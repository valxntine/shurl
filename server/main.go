package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/valxntine/shurl/internal/repository"
	"github.com/valxntine/shurl/internal/rest"
	"github.com/valxntine/shurl/internal/service"
)

func main() {
	PORT := os.Getenv("PORT")

	// urlRepo := repository.NewInMemoryRepository()
	urlRepo, err := repository.NewRepository()
	if err != nil {
		log.Fatalln(err)
	}

	urlService := service.NewUrlService(urlRepo, urlRepo)

	r := rest.NewRouter(rest.RouterConfig{
		BusinessLogic: rest.Services{
			ShortUrlCreater:   &urlService,
			OriginalUrlGetter: &urlService,
		},
	})

	server := http.Server{
		ReadHeaderTimeout: time.Second * 5,
		Addr:              fmt.Sprintf(":%s", PORT),
		Handler:           r,
	}

	log.Println(fmt.Sprintf("service started on port [::] %s", server.Addr))
	log.Fatal(server.ListenAndServe())
}
