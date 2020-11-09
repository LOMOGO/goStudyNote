package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	//采用这种方式或者
	/*r := gin.Default()

	http.ListenAndServe(":8080", r)*/

	r := gin.Default()

	s := &http.Server{
		Addr:         "8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
