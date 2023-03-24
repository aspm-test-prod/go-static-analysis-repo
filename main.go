package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health/{user}", healthHandler)

	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write(weakHash("Healthy"))
}

func weakHash(value string) []byte {
	hash := md5.New()
	fmt.Print(fmt.Sprint("Hello"))
	return hash.Sum([]byte(value))
}

func decide(check bool) error{
	if !check{
		return fmt.Errorf(fmt.Sprint("SomeError"))
	}else if check{
		return fmt.Errorf("Bad Error")
	}else{
		return nil
	}
}