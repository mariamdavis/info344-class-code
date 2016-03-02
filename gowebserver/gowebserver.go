package main

import (
    "net/http"
    "fmt"
    "time"
    "encoding/json"
    "log"
    "runtime"
)

//HelloResponse represents a response from the hello route
type HelloResponse struct{
    Name string `json:"name"`
    Message string 
    GeneratedAt time.Time
}

var memstats = new(runtime.MemStats)

func getMemStats(w http.ResponseWriter, r *http.Request) {
    runtime.ReadMemStats(memstats)
    allocstats := make(map[string]uint64)
    allocstats["alloc"] = memstats.Alloc
    allocstats["totalAlloc"] = memstats.TotalAlloc
    j, err := json.Marshal(allocstats)
    if nil != err {
       log.Println(err)
       w.WriteHeader(500)
       w.Write([]byte(err.Error())) 
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
    }
}

//* = pointer to http request object
func sayHello(w http.ResponseWriter, r *http.Request){
    //take url path and give slice that starts path char pattern. give me whats beyond it
    name := r.URL.Path[len("/api/vi/hello/"):]
    resp := HelloResponse{Name: name, 
                          Message:"Hello " + name,
                          GeneratedAt: time.Now()}
    j, err := json.Marshal(resp) //serialization
    if nil != err {
        log.Println(err)
        w.WriteHeader(500)
        w.Write([]byte(err.Error()))
    } else {
        w.Header().Add("Content-Type", "application/json")
        w.Write(j)
    }
    
   // w.Write([]byte("Hello " + name))
} 


func main() {
    //http.HandleFunc("/", sayHello)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/api/v1/hello/", sayHello) // /=means anything after it will call the same function
    http.HandleFunc("/api/v1/memstats", getMemStats)
    fmt.Println("Server listening on port 9000...")
    http.ListenAndServe(":9000", nil) 
}

