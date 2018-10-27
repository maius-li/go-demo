package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Path: ", r.RequestURI)
		log.Println("Method: ", r.Method)

		// GET only Query
		log.Println("Query: ", r.URL.Query())
		log.Println("Query a: ", r.URL.Query().Get("a"))

		// Post has form, body
		// body, err := ioutil.ReadAll(r.Body)
		// log.Println("Read Error: ", err)
		// log.Println("Body: ", string(body))
		// body := []byte(`{
		// 	"Name":"Maius"
		// }`)

		var data person
		err := json.NewDecoder(r.Body).Decode(&data)
		// err := json.Unmarshal(body, &data)
		log.Println("JSON Error: ", err)
		log.Println("Person: ", data.Name)
		log.Println("Form Z: ", r.FormValue("z"))

		w.Write([]byte(`hello world`))
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// a := []int{8, 5, 1, 1}
		// b := map[string]string{"name": "zoular", "age": "35"}
		// var b map[string]string
		// b = make(map[string]string)
		age := 20
		var b *person
		// b := &person{}
		if age > 20 {
			b = &person{Name: "Zuolar"}
		} else {
			b = &person{Name: "Maius"}
		}

		// jsonData, err := json.Marshal(b)
		w.Header().Set("Content-Type", "application/json")
		// w.Write(jsonData)
		err := json.NewEncoder(w).Encode(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3000", nil)
}

type person struct {
	Name string
}
