package main

import "fmt"
import "time"

func worker1( c chan int ){
    for{
        time.Sleep( 2 * time.Second )
        c <- 10
    }
}

func worker2( c chan int ){
    for{
        time.Sleep( 4 * time.Second )
        c <- 20
    }
}

func timer( stop chan int ){
    time.Sleep( 5 * time.Second )
    stop <- 1
}

func main(){
    c1 := make( chan int )
    c2 := make( chan int )
    c3 := make( chan int )
    go worker1( c1 )
    go worker2( c2 )
    // select coammand will wait untill data have come to varaible can check
    stop := make( chan int)
LOOP: // Label line to use for break
    for{
        select { // command to make we can wait data from to many channal
            case n := <- c1:
                fmt.Println( n )
            case n := <- c2:
                fmt.Println( n )
            case c3 <- 10:
                fmt.Println( "send to c3")
            case <- stop:
                break LOOP  // have declare label because if don't declare will break select
            default: // will run in case don't have to check
                time.Sleep( 1 * time.Second) // use this for sleep wait
                fmt.Println(".")
        }
    }
    // warning select case will do if can do because select will check always
}
