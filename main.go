package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "location to you config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handle error
		fmt.Println("error")

	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("error")
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("error:getpodlist")

	}
	for i := 0; i < len(pods); i++ {
		fmt.Println("OK", pods[i])

	}
}
