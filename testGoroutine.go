package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go Print(&wg)
	wg.Wait()

}
func Print(wg *sync.WaitGroup){
	fmt.Println("aaaaa")
	wg.Done()
}
