package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getTop100(c *gin.Context) {

	/* Make REST call to EPSS API here */
	response, err := http.Get("https://api.first.org/data/v1/epss")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func main() {
	router := gin.Default()
	router.GET("/epss", getTop100)

	router.Run("localhost:8080")
}
