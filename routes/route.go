package routes

/*
This class is a route handler. From here, the requests are directed towards the controller
*/

import (
	"github.com/gin-gonic/gin"
	"takeHomeTest/controllers"
)

type Route interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller controllers.Controller
}

func RegisterHandlers(engine *gin.Engine, controller controllers.Controller) Route {
	return &route{
		engine:     engine,
		controller: controller,
	}
}

//This is a route handler for various requests
func (r route) RegisterHandlers() {
	r.engine.POST("/createPayout", r.controller.CreatePayout)
}
