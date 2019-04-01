package main

import "fmt"
import "time"

func worker1( c chan int ){
    time.Sleep( 1 * time.Second )
    c <- 10
}

func worker2( c chan int ){
    time.Sleep( 4 * time.Second )
    c <- 20
}

func main(){
    c1 := make( chan int )
    c2 := make( chan int )
    go worker1( c1 )
    go worker2( c2 )
    select { // command to make we can wait data from to many channal
        case n := <- c1:
            fmt.Println( n )
        case n := <- c2:
            fmt.Println( n )
    }
}
