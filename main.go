package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tylerb/graceful"

	"github.com/kyokomi/example_goa_api/application/controllers"
	"github.com/kyokomi/example_goa_api/gen/app"
)

type mode string

const (
	developMode    mode = "dev"
	productionMode mode = "prd"
)

func parseMode(modeStr string) mode {
	return mode(modeStr)
}

func (m mode) isProduction() bool {
	return m == productionMode
}

func (m mode) isDevelop() bool {
	return m == developMode
}

func newGoaService(m mode) *goa.Service {
	// Create service
	service := goa.New("example_goa_api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(m.isDevelop()))
	service.Use(middleware.ErrorHandler(service, m.isDevelop()))
	service.Use(middleware.Recover())

	app.MountHealthController(service, controllers.NewHealthController(service))
	if m.isDevelop() {
		app.MountSwaggerController(service, controllers.NewSwaggerController(service))
	}
	return service
}

func serve(modeStr string, addr string, timeoutSec int) {
	m := parseMode(modeStr)
	service := newGoaService(m)

	// Setup graceful server
	server := &graceful.Server{
		Timeout: time.Duration(timeoutSec) * time.Second,
		Server:  &http.Server{Addr: addr, Handler: service.Mux},
	}
	// And run it
	log.Printf("ListenAndServe %s mode [%s]", addr, m)
	log.Print(server.ListenAndServe())
}

func main() {
	addr := os.Getenv("PORT")
	if len(addr) == 0 {
		flag.StringVar(&addr, "p", ":8080", "server run port (default :8080) ")
	}
	timeoutSec := flag.Int("t", 15, "time (sec) to request timeout (default 15) ")
	mode := flag.String("m", "prd", "server run mode (default prd) ")
	flag.Parse()

	serve(*mode, addr, *timeoutSec)
}
