package cube

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
)

type Application struct {
	Conf *Conf
}

func (app *Application) initConfig(configFile string) {
	app.Conf = &Conf{}
	if _, err := toml.DecodeFile("./config/application.toml", &app.Conf.ApplicationConf); err != nil {
		panic(err)
	}
	fmt.Println("init --->")
	fmt.Println(configFile)
	fmt.Println(app.Conf)
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
