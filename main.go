package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mateusprt/auth-api/src/handlers"
	"github.com/subosito/gotenv"
)

func init() {
	/*key := make([]byte, 64)
	_, err := rand.Read(key)

	if err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)

	fmt.Println(stringBase64)*/
	gotenv.Load()
}

func main() {

	routes := handlers.LoadAll()
	port := ":" + os.Getenv("PORT")

	server := &http.Server{
		Addr:    port,
		Handler: routes,
	}

	log.Println("Server is running on port " + port + "...")
	err := server.ListenAndServe()

	if err != nil {
		log.Println("Server failed")
		log.Panic(err)
	}

}
