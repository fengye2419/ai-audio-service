package routers

import (
	"context"

	"github.com/go-chi/chi/v5" // Add this line to import the middleware package if needed
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	apiv1 "gl.eeo.im/fengye2419/ai-audio-service/internal/app/routers/api/v1"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/app/routers/common"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/app/routers/web"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/setting"
	"gl.eeo.im/fengye2419/ai-audio-service/internal/pkg/util"
)

// GlobalInit initializes the global settings
func GlobalInit() {
	setting.NewService()
	// init db engine
	if err := common.InitDBEngine(context.Background()); err != nil {
		logrus.Fatalf("Failed to initialize ORM engine: %v", err)
	} else {
		logrus.Infof("ORM engine initialized.")
	}
}

// NormalRoutes returns the normal routes
func NormalRoutes() *chi.Mux {
	r := chi.NewRouter()
	// request id middleware
	r.Use(util.RequestID)
	// middlewares
	for _, middle := range common.Middlewares() {
		r.Use(middle)
	}
	// web routes
	r.Mount("/", web.Routes())
	// api v1
	r.Mount("/api/v1", apiv1.Routes())
	// swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
