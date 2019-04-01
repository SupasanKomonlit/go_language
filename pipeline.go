package main

import "fmt"

func generator() chan int{
    output := make( chan int )
    go func(){
       for i := 0; i < 10; i++{
            output <- i
        }
        close( output )
    }()
    return output
}

func multiplier( input chan int ) chan int{
    output := make( chan int)
    go func(){
        for num := range input{ // if you want to use range you must to have close( channel )
            output <- num*num
        }
        close( output )
    }()
    return output
}

func adder( input chan int ) chan int{
    output := make( chan int )
    go func(){
        for num := range input{
            output <- num+num
        }
        close( output )
    }()
    return output
}

func print( input chan int ){
    for num := range input{
        fmt.Println( num )
    }
}

func main(){
    g := generator()
    m := multiplier(g)
    a := adder( m )
    print( a )
}
