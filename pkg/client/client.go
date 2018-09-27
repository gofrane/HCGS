package client



import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	//  Leave blank for the default context in your kube config.
	context = ""


)



func Client() ( *kubernetes.Clientset , error  ){
	//  Get the local kube config.
	fmt.Printf("Connecting to Kubernetes Context %v\n", context)
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: context}).ClientConfig()
	if err != nil {
		panic(err.Error())
	}

	// Creates the clientset
	clientset, err1 := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err1.Error())
	}

	//  Scale our replication controller.
	fmt.Printf("Scaling replication controller ", clientset)


	return  clientset,err1
}
