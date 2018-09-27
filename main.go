package main

import (
	"HCGS/pkg/client"
	"fmt"
)

func  main()  {



clientset , error:=client.Client()

if error!=nil {
	
	
	fmt.Println("error of client ")
	
}else {
	
	fmt.Println("clienset", clientset)
}


}
