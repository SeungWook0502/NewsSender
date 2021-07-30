package main
 
import (
    "fmt"
)
 
func main() {

    var a [3]int
    a[0] = 1
    a[1] = 2
    a[2] = 3

    fmt.Println(a[0],a[1],a[2])

    aa(&a[2])

    fmt.Println(a[2])

    b := make([]string,0,2)
    fmt.Println(len(b))
    b=append(b,"aa")
    fmt.Println(len(b))
    b=append(b,"bb")
    fmt.Println(len(b))
    fmt.Println(b)

}

func aa(s *int){
    
    *s = 3

    fmt.Println(*s)

}