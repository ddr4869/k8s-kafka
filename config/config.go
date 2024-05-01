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
	DB        DBConf
	K8sClient *kubernetes.Clientset
}

type GinConf struct {
	Mode string
}

type DBConf struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
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
		DB: DBConf{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		},
		K8sClient: clientset,
	}
}
