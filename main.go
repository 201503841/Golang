package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"html/template"
	"net/http"
)

//Respuesta is
type Respuesta struct {
	Analisis string `json:"analisis"`
	Grafo    string `json:"grafo"`
	Errores  string `json:"errores"`
}

//Codigo is
type Codigo struct {
	Codigo string `json:"codigo"`
}

func getTraduccion(w http.ResponseWriter, r *http.Request) {
	var url = "http://localhost:3000/analisis/"
	var decoder = json.NewDecoder(r.Body)
	var c Codigo
	err := decoder.Decode(&c)

	if err != nil {
		panic(err)
	}
	var jsonStr = []byte(`{"codigo":"` + c.Codigo + `"}`)
	fmt.Println(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyBytes))
	fmt.Fprintf(w, string(bodyBytes))
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, nil)
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/traducir", getTraduccion)
	fmt.Printf("Servidor escuchando en: http://localhost:8002/")
	http.ListenAndServe(":8002", nil)
}
