package main

import (
	// Utilities "github.com/they-them/slash-wiki-rest-go/cmd/slash-wiki-rest-go/utilities" 
	Structures "github.com/they-them/slash-wiki-rest-go/cmd/slash-wiki-rest-go/structures" 
	"github.com/gorilla/schema"
	"github.com/gorilla/mux"
	"encoding/json"
	"time"
	"net/http"
	"fmt"
	"os"
)

var decoder = schema.NewDecoder()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		
		if nil != err {
			// Handle error
		}

		var parameters Structures.Parameters
		err = decoder.Decode(&parameters, request.PostForm)

		if nil != err {
			// Handle error
		}

		message := Structures.Message{
			ResponseType: "in_channel",
			Text:         fmt.Sprintf("Hello %s, the time is %s.", parameters.UserId, time.Now().Format(time.RFC850)),
		}

		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(message)
	})

	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}