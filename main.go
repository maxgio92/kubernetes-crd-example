package main

import (
	"flag"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"github.com/maxgio92/kubernetes-crd-example/api/types/v1alpha1"
	clientV1alpha1 "github.com/maxgio92/kubernetes-crd-example/clientset/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to kubeconfig file")
	flag.Parse()
}

func main() {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		conifg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err)
	}

	clientset, err := clientv1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	projects, err := clientSet.Projects("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("projects found: %v\n", projects)
}
