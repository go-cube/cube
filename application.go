package cube

import (
	"fmt"
	"net/http"
	"time"
)

type Application struct {
}

func (app *Application) initConfig(configFile string) {
	fmt.Println("init --->")
	fmt.Println(configFile)
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
}

func (app *Application) Start() {

	fmt.Printf("Starting application on address `%s`\n", ":8080")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
