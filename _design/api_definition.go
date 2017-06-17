package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("example_goa_api", func() {
	Title("The example goa api")
	Description("A basic example of a api with goa")
	Host("localhost:8081")
	Scheme("http")
	BasePath("/")

	Origin("http://swagger.example.dev", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})
})
