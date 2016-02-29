package main

import (
    "fmt"
)

func doubleIt(val *int){
    //val *= 2
    *val *= 2 //creates pointer
}

func main() {
    val := 5
   // doubleIt(val)
    doubleIt(&val);
    fmt.Println(val)
}