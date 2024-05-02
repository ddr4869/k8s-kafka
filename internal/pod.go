package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) GetAllPods(c *gin.Context) {
	pods, _ := s.K8sClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	podList, err := dto.V1PodListToJson(pods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println(podList)
	c.JSON(http.StatusOK, podList)
}

func (s *Server) GetNamespacePods(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetNamespacePodRequest)
	v1Pods, _ := s.K8sClient.CoreV1().Pods(reqUri.Namespace).List(context.TODO(), metav1.ListOptions{})
	podList, err := dto.V1PodListToJson(v1Pods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, dto.NewSuccessResponse(
		http.StatusOK, "success", podList,
	))
}

func (s *Server) GetPod(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetNamespacePodRequest)
	req := c.MustGet("req").(dto.GetPodNameRequest)
	v1Pod, _ := s.K8sClient.CoreV1().Pods(reqUri.Namespace).Get(context.TODO(), req.Name, metav1.GetOptions{})
	pod, err := dto.V1PodToJson(v1Pod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, pod)
}

func (s *Server) GetPodLog(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetNamespacePodRequest)
	req := c.MustGet("req").(dto.GetPodNameRequest)
	podLog, err := s.K8sClient.CoreV1().Pods(reqUri.Namespace).GetLogs(req.Name, &v1.PodLogOptions{}).Stream(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer func() {
		fmt.Println("close podLog")
		podLog.Close()
	}()
	p := make([]byte, 1024)
	for {
		n, err := podLog.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}
	c.JSON(http.StatusOK, "success")
}

func (s *Server) CreateNamespacePods(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateNamespacePodRequest)
	yamlFile, err := os.ReadFile(s.config.KubeYaml + req.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	v1Pod := &v1.Pod{}
	if err := v1Pod.Unmarshal(yamlFile); err != nil {
		panic(err.Error())
	}
	_, err = s.K8sClient.CoreV1().Pods("default").Create(context.TODO(), v1Pod, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "success!")
}
