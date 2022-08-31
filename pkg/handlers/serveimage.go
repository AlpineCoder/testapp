package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/AlpineCoder/testapp/pkg/generators"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {

	var fileBytes []byte

	files, err := ioutil.ReadDir("images")
	if err != nil {
		log.Print(err)
		fileBytes, err = ioutil.ReadFile(generators.GenRandomImage())
		if err != nil {
			log.Print("oh no!")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
		return
	}

	imageIndex := rand.Intn(len(files) - 1)

	fmt.Println("http://" + r.Host + "/images/" + files[imageIndex].Name())

	fileBytes, err = ioutil.ReadFile("images/" + files[imageIndex].Name())
	if err != nil {
		log.Print("oh no!")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
