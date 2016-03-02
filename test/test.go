package main

import (
    "net/http"
    "os"
    "log"
    "encoding/csv"
    "io"

)
func loadData(){
  csvfile, err := os.Open("data/wordsEn.txt")
  if err != nil {
                 log.Fatal(err)
                 return
         }
  reader := csv.NewReader(csvfile)
  for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
        fmt.Println(record)
	}
    
     defer csvfile.Close()
}

func main() {
    loadData()
    fmt.Println("Server listening on port 9000...")
    http.ListenAndServe(":9000", nil) 
}

