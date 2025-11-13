package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		fmt.Printf("%s", string(b))

		reader := bytes.NewReader(b)
		res, err := http.Post("https://dev.raiffcycles.ch/api/linka", "application/json", reader)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(body))
	})

	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.config")

	viper.SetConfigName("unchunking-proxy")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("UNCHUNKING")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetDefault("port", 8080) // 10 minutes

	viper.SafeWriteConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}
