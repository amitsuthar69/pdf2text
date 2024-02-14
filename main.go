package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	_500 := http.StatusInternalServerError
	_400 := http.StatusBadRequest

	r.Post("/convert", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("pdf")
		if err != nil {
			http.Error(w, "Error reading file", _400)
			return
		}
		defer file.Close()

		// Buffer for storing uploaded PDF data
		buf := bytes.NewBuffer(nil)

		// Copy uploaded data to buffer
		_, err = io.Copy(buf, file)
		if err != nil {
			http.Error(w, "Error copying file data", _500)
			return
		}

		// Create command with pipes
		cmd := exec.Command("pdftotext", "-", "-")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			http.Error(w, "Error creating stdin pipe", _500)
			return
		}
		defer stdin.Close()

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			http.Error(w, "Error creating stdout pipe", _500)
			return
		}
		defer stdout.Close()

		// Start the command
		err = cmd.Start()
		if err != nil {
			http.Error(w, "Error starting pdftotext", _500)
			return
		}

		// Write uploaded data to stdin
		_, err = stdin.Write(buf.Bytes())
		if err != nil {
			http.Error(w, "Error writing to pdftotext stdin", _500)
			return
		}
		stdin.Close()

		// Read extracted text from stdout
		text, err := io.ReadAll(stdout)
		if err != nil {
			http.Error(w, "Error reading extracted text", _500)
			return
		}

		//? Replacing newlines with spaces -- can degrade performance for larger documents!
		text = []byte(strings.ReplaceAll(string(text), "\n", " "))

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(struct {
			Text string `json:"text"`
		}{
			Text: string(text),
		})
		if err != nil {
			http.Error(w, "Error encoding response", _500)
			return
		}
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

//* test the api := curl -X POST -F "pdf=@sample.pdf" http://localhost:8080/convert > output.txt
