package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	Gin       GinConf
	K8sClient *kubernetes.Clientset
}

type GinConf struct {
	Mode string
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/ieungyu/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)

	return &Config{
		Gin: GinConf{
			Mode: os.Getenv("GIN_MODE"),
		},
		K8sClient: clientset,
	}
}
