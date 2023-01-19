package main

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	flag.String("kubeconfig", "~/.kube/config", "location to you config file")
	config, err := clientcmd.BuildConfigFromlogs("", kubeconfig)
	if err != nil {
		//handle error

	}
	fmt.Println(config)
}
