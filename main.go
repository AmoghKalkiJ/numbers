package main

import (
	"fmt"
	"time"
)

var i int


func even (done chan bool){
	for{
		select {
		case <-done:
			    i++
				fmt.Println("in even",i)
		        done <- true

		default:
			fmt.Println("in default",i)
			done<-true
			}
	}
}

func odd (done chan bool) {
	for {

		select {
		case <-done:
			i++
			fmt.Println("in odd",i)
			done <- true
			/*case i==num:
				fmt.Println(i)
				fmt.Println("done printing" )
				ch<-true	*/
		}
	}
}



func main(){
	done:=make(chan bool)
	go even (done)
	go odd (done)
	time.Sleep(2*time.Second)
}
