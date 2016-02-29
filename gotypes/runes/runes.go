package main

import (
    "fmt"
)

func main() {
    s:= "Hello World!"
    
    //idx gives back index
    for idx, r := range s {
       fmt.Println(idx, string(r))
    }
    
}