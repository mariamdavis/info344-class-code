package main

import (
    "github.com/mariamdavis/info344-class-code/gotypes/structs/person"
)

func main() {
    prs := person.NewPerson("Mariam", "Davis")

    //person.SayHello(prs) 
    prs.SayHello()
}