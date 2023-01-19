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
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("error:getpodlist")

	}
	fmt.Println("---Pods from default namespace---")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)

	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	fmt.Println("---Deployments from default namespace---")
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)

	}
	NamespaceList, err := clientset.CoreV1().Namespaces("").List(context.TODO(), metav1.ListOptions{})
	for _, n := range NamespaceList.Items {
		fmt.Println(n.Name)
	}

}
