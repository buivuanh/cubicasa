package main

import (
	"github.com/buivuanh/cubicasa/api"
	"github.com/buivuanh/cubicasa/app"
	"github.com/buivuanh/cubicasa/model"
	"github.com/go-playground/kms"
	"github.com/go-playground/kms/kmsnet/kmshttp"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", "5000")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	conf := model.AppSettings{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	a, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	API, err := api.New(a, e)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", API.Routes.Root))

	kms.ListenTimeout(false, 10*time.Second)

	// Start server
	log.Println("Server started")
	err = kmshttp.ListenAndServe(":"+viper.GetString("PORT"), mux)
	if err != nil {
		log.Fatal(err)
	}
}
