package main

import "fmt"

func count( c chan int ){
    for i := 0 ; i < 20 ; i++ {
        c <- i
    }
    close( c )
}

func main(){
    c := make( chan int)
    for i := range c{
        fmt.Println(i)
    }
}
