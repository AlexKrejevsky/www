package main

import ("fmt"; "net/http"; "html/template")

type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grade, happiness float64
}

func home_page(w http.ResponseWriter, r *http.Request)  {
  bob := User{"Bob", 25, -50, 4.2, 0.8}
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {
fmt.Fprintf(w, "Alex phone 88005553535")

}

func handlerRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}
func main() {
  //var bob User
  handlerRequest()
}
