package main 

import(
    "fmt"
    "time"
)
//use slice when dynamic

func main () {
  var months [13]string
  for idx := 0; idx < 12; idx++ {
      months[idx] = time.Month(idx+1).String()//wants index 1 - 12 not 0
  }
  fmt.Println(months)
  fmt.Println(months[1:2])
  fmt.Println(months[5])
  fmt.Println(months[5:]) //grabs slice of array
  
  var dynoMonths []string 
  
  for idx := 0; idx < 12; idx++ {
    dynoMonths = append(dynoMonths, time.Month(idx+1).String())
   }   
     fmt.Println(dynoMonths)
     
    
}