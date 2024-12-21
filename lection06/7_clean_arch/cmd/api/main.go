package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guap-courses/go/lection06/7_clean_arch/internal/handlers"
	"github.com/guap-courses/go/lection06/7_clean_arch/internal/repository"
	"github.com/guap-courses/go/lection06/7_clean_arch/internal/services"
	pkglog "github.com/guap-courses/go/lection06/7_clean_arch/pkg/log"
	pkgpostgres "github.com/guap-courses/go/lection06/7_clean_arch/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	dsn := "postgres://username:password@localhost:5442/guap_course?sslmode=disable"

	pool, err := pkgpostgres.NewPool(dsn, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool, logger)

	candlesService := services.NewCandlesService(logger, repo)

	r := chi.NewRouter()
	r.Use(pkglog.NewStructuredLogger(logger))

	candlesHandler := handlers.NewCandles(logger, candlesService)
	r.Mount("/candles", candlesHandler.Routes())

	http.ListenAndServe(":3000", r)
}
