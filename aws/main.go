package main

import ("fmt"
        "net/http")

func home_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "Hello form GitHub and AWS!!")
}

func main(){
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080",nil)
  }
