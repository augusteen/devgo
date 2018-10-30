package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<title>Hello</title>
			</html>

		`))
	})

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("ListenandServe:", err)
	}
}
