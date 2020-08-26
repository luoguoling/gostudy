package main

import (
	"fmt"
	"sync"
)

func main(){
	c1 := make(chan int)
	var readc <-chan int  = c1
	var writec chan <- int = c1
	c2 := make(chan int)
	var readc2 <-chan int  = c2
	var writec2 chan <- int = c2
	var wg sync.WaitGroup
	wg.Add(3)
	go setChan(writec,&wg)
	go setChan2(writec2,&wg)

	go getChan(readc,&wg)
	go getChan2(readc2,&wg)
	wg.Wait()
}
func setChan(writec chan <- int,wg *sync.WaitGroup){
	for i:=0;i<10;i++{
		writec <- i
	}
	wg.Done()
}
func setChan2(writec2 chan <- int,wg *sync.WaitGroup){
	for i:=10;i<20;i++{
		writec2 <- i
	}
	wg.Done()
}
func getChan(readc <-chan int,wg *sync.WaitGroup)  {
	for i:=0;i<10;i++{
		fmt.Printf("取出来的值是%d\n",<-readc)

	}
	wg.Done()

}
func getChan2(readc2 <-chan int,wg *sync.WaitGroup)  {
	for i:=0;i<10;i++{
		fmt.Printf("AAAA取出来的值是%d\n",<-readc2)

	}
	wg.Done()

}
