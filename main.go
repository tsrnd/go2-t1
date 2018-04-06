package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/router"
	"github.com/tsrnd/trainning/shared/monitoring"
)

func main() {
	// start internal server
	go startInt()

	// start external server
	startExt()
}

func startExt() {
	// sql new.
	sqlHandler := infrastructure.NewSQL()
	// s3 new.
	s3Handler := infrastructure.NewS3()
	// cache new.
	cacheHandler := infrastructure.NewCache()
	// logger new.
	loggerHandler := infrastructure.NewLogger()

	// monitoring setup
	mLogger := infrastructure.NewLoggerWithType("monitoring")
	monitoring.Setup(mLogger)

	mux := chi.NewRouter()
	r := &router.Router{
		Mux:           mux,
		SQLHandler:    sqlHandler,
		S3Handler:     s3Handler,
		CacheHandler:  cacheHandler,
		LoggerHandler: loggerHandler,
	}

	r.InitializeRouter()
	r.SetupHandler()

	// after process
	defer infrastructure.CloseLogger(r.LoggerHandler.Logfile)
	defer infrastructure.CloseRedis(r.CacheHandler.Conn)
	defer infrastructure.CloseLogger(mLogger.Logfile)

	_ = http.ListenAndServe(":8080", mux)
}

func startInt() {
	mux := chi.NewRouter()
	logger := infrastructure.NewLogger()
	// mux.Use(mMiddleware.Logger(logger))

	defer infrastructure.CloseLogger(logger.Logfile)

	// profile
	mux.Mount("/debug", middleware.Profiler())
	_ = http.ListenAndServe(":18080", mux)
}
