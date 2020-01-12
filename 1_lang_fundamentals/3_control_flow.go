package main

import "fmt"

func main() {

    //traditional for loop
    // for i := 0; i < 10; i++ {
    //  for j := 0; j <= i; j++ {
    //      fmt.Println(i, j)
    //  }
    // }

    //while for in go
    i := 0
    for i < 10 {
        i += 1
        if i == 5 {
            continue
        }
        fmt.Printf("%d\t", i)
    }

    fmt.Println()
    y := 0
    for {
        if y > 9 {
            break
        }
        fmt.Println(y)
        y++
    }

    //if with declaration -> if declaration; condition{ code }
    if x := 2; x == 2 {
        fmt.Println("True")
    }

    //if elseif and else
    n := 32
    if n == 32 {
        fmt.Println("its 32")
    } else if n == 40 {
        fmt.Println("its 40")
    } else {
        fmt.Println("I have no idea")
    }

    sw := 6
    switch sw{
        case 1:
            fmt.Println("1")
        case 2:
            fmt.Println("2")
        case 3:
            fmt.Println("3")
        default:
            fmt.Println("I have no Idea")
    }
}
