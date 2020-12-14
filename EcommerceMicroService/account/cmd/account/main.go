package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ausaafnabi/microservice-graphql/EcommerceMicroService/retry"
	"github.com/ausaafnabi/microservice-graphql/EcommerceMicroService/account"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}