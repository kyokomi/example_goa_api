package controllers

import (
	"github.com/goadesign/goa"
	"github.com/kyokomi/example_goa_api/gen/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// TODO: check db connection etc...
	return ctx.OK([]byte("ok"))
}
