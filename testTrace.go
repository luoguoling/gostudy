package main

import (
	"fmt"
	"os"
	"runtime/trace"

)
import "sync"
func mockSendToServer(url string){
	fmt.Printf("SERVER URL:%s \n",url)
}
func main()  {
	f,err := os.Create("trace2.out")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil{
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:1000","0.0.0.0:2000","0.0.0.0:3000","0.0.0.0:4000"}
	wg := sync.WaitGroup{}
	for _,url := range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			mockSendToServer(url)
		}(url)

	}
	wg.Wait()
}
