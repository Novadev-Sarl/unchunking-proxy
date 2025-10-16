package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		fmt.Printf("%s", string(b))

		reader := bytes.NewReader(b)
		res, err := http.Post("https://dev.raiffcycles.ch/api/linka", "application/json", reader)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(body))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
