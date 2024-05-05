package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ddr4869/k8s-kafka/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 모든 요청을 허용합니다. 보안상의 이유로 변경해야 할 수 있습니다.
	},
}

func (s *Server) GetAllPods(c *gin.Context) {
	pods, _ := s.K8sClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	podList, err := dto.V1PodListToJson(pods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
		return
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
		return
	}
	c.JSON(http.StatusOK, pod)
}

func (s *Server) GetPodLog(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetNamespacePodRequest)
	req := c.MustGet("req").(dto.GetPodNameRequest)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	podLogStream, err := s.K8sClient.CoreV1().Pods(reqUri.Namespace).GetLogs(req.Name, &v1.PodLogOptions{}).Stream(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer podLogStream.Close()

	// 로그를 지속적으로 수신하고 웹소켓을 통해 클라이언트에게 전송합니다.
	go func() {
		defer fmt.Println("closing podLog stream")
		buf := make([]byte, 1024)
		for {
			// Pod의 로그를 읽어옵니다.
			n, err := podLogStream.Read(buf)
			if err != nil {
				if err != io.EOF {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				// 에러가 EOF일 경우 무시하고 계속 수신합니다.
				continue
			}
			// 로그 데이터를 WebSocket을 통해 클라이언트에게 전송합니다.
			if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}()

	// // go func() {
	// defer fmt.Println("closing podLog stream")
	// buf := make([]byte, 1024)
	// for {
	// 	n, err := podLog.Read(buf)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 	}
	// 	fmt.Println(4)
	// 	// Send log data over WebSocket
	// 	if err := conn.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	fmt.Println(5)
	// }
	// }()

	c.JSON(http.StatusOK, "streaming start")
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

func (s *Server) SocketHandler(c *gin.Context) {
	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer conn.Close()
	conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			log.Printf("conn.ReadMessage: %v", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("conn.WriteMessage: %v", err)
			return
		}
	}
}
