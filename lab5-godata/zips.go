package main

import (
    //"fmt"
    "os"
    "log"
    "encoding/csv"
    "io"
   
)

type records struct{
    Total int
    RecordsPerDistinct int
    HighestPop string
    
}


func main() {
  csvfile, err := os.Open("zip_code_database.csv")
  if err != nil {
                 log.Fatal(err)
                 return
         }
  reader := csv.NewReader(csvfile)
  record, err := reader.Read()
  pop := 0
  for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
         i := record[14]

		//fmt.Println(i);
	}
     fmt.Println(pop)
     defer csvfile.Close()
}