package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestMatchesCreate(t *testing.T) {
	r := gofight.New()

	r.POST("/matches/create").
		Run(newRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
func TestMatchesJoin(t *testing.T) {
	r := gofight.New()

	r.POST("/matches/1/join").
		Run(newRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusAccepted, r.Code)
		})
}

func TestMatchesStatus(t *testing.T) {
	r := gofight.New()

	r.GET("/matches/1/status").
		Run(newRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestMatchesMove(t *testing.T) {
	r := gofight.New()

	r.PUT("/matches/1/move").
		Run(newRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
