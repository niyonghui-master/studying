package main

import "fmt"

type People struct {
   Name string `json:"name"`
}

func main()  {
   js := `{
   "name":"11"
   }
   `
var p People
   err := jsoniter.Unmarshal([]byte(js),&p)
   if err !=nil {
      fmt.Println("err:",err)
      return
   }
   fmt.Println("people:",p)
}