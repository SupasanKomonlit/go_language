package main

import "fmt"

func worker( buff chan int ){
    for data := range buff{
        fmt.Println( data )
    }
}

func main(){
    bufferChannel := make( chan int , 3 )
    bufferChannel <- 10
    bufferChannel <- 11
    bufferChannel <- 12
    close( bufferChannel )
    worker( bufferChannel )
}
