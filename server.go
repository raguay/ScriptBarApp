package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type DataType struct {
	Days string `json:"days"`
	Age  string `json:"age"`
	Text string `json:"text"`
	IP   string `json:"ip"`
}

type VariableType struct {
	Name  string   `json:"name" binding:"required"`
	Value DataType `json:"value" bindings:"required"`
}

func RunServer(a *App, ctx context.Context) {
	//
	// This will have the web server backend for BulletinBoard.
	//
	r := gin.Default()
	r.Use(gin.Recovery())

	//
	// Define the message route. The message is given on the URI string and in the body.
	//
	r.GET("/api/getvar/:variable", func(c *gin.Context) {
		variable := c.Param("variable")

		//
		// Send the request
		//
		rt.EventsEmit(ctx, "getvariable", variable)

		//
		// Get the return.
		//
		running := true
		rt.EventsOn(ctx, "returnvariable", func(optionalData ...interface{}) {
			c.JSON(http.StatusOK, optionalData)
			running = false
			rt.EventsOff(ctx, "returnvariable")
		})
		for running {
			time.Sleep(time.Millisecond)
		}
	})

	//
	// Add route for listing the variables.
	//
	r.GET("/api/getvar/list", func(c *gin.Context) {
		//
		// Send the request
		//
		rt.EventsEmit(ctx, "listvariables", nil)

		//
		// Get the return.
		//
		running := true
		rt.EventsOn(ctx, "returnvariablelist", func(optionalData ...interface{}) {
			c.JSON(http.StatusOK, optionalData)
			running = false
			rt.EventsOff(ctx, "returnvariablelist")
		})
		for running {
			time.Sleep(time.Millisecond)
		}
	})

	//
	// Add the route for setting a variable value.
	//
	r.PUT("/api/setvar/:variable", func(c *gin.Context) {
		var json VariableType
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//
		// Send it to the frontend.
		//
		rt.EventsEmit(ctx, "setvariable", json)

		c.JSON(http.StatusOK, "okay")
	})

	//
	// Run the server.
	//
	r.Run(":9696")
}
