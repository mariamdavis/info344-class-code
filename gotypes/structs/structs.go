package main

import (
    "github.com/mariamdavis/info344-class-code/gotypes/structs/person"
    "fmt"
)

func main() {
    prs := person.NewPerson("Mariam", "Davis")

    //person.SayHello(prs) 
    fmt.SayHello()
    prs.SayHello()
}