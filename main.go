package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	board [][]int

	move struct {
		Position [2]int `json:"pos"`    // [x,y]
		Action   string `json:"action"` //add, remove, pass, forfeit?
		Player   bool   `json:"player"`
	}

	match struct {
		// ID string `json:"id"` //already the key on the matches
		Players [2]int  `json:"player_ids"`
		Board   board   `json:"board"`
		History []move  `json:"history"`
		Score   [2]uint `json:"score"`
	}

	status struct {
		State int   `json:"state"`
		Board board `json:"board"`
		Last  bool  `json:"last_player"`
	}
)

var matches map[string]match

// to-do:
// 	[] game logic
func main() {
	e := newRouter()
	e.Logger.Fatal(e.Start(":1323"))
}

func newRouter() *echo.Echo {
	e := echo.New()
	e.POST("/matches/create", matchesCreate)
	e.POST("/matches/:id/join", matchesJoin)
	e.GET("/matches/:id/status", matchesStatus)
	e.PUT("/matches/:id/move", matchesMove)
	return e
}

func matchesCreate(c echo.Context) error {
	// to-do: create good, readable id, create the match, register first user
	return c.String(http.StatusOK, "{id:\"666fff\"}")
}

func matchesJoin(c echo.Context) error {
	// to-do: register new user, add it to the list and change match
	// state (cookies?)
	if true /*actual check to see if the user is allowed to join*/ {
		return c.String(http.StatusAccepted, "join"+c.Param("id"))
	}
	return c.String(http.StatusNoContent, "join"+c.Param("id"))
	//https://godoc.org/net/http#pkg-constants
}

func matchesStatus(c echo.Context) error {
	/*
		to-do: return match status {
			state:waiting_opponent/playing/over
			board:[
				[0,0,0,1,2],
				[0,0,0,1,2],
			]
			last_to_play:0/1,
			diff:[
				{
					type:add/rem,
					pos:[x,y],
					player:0/1
				},
				...
			]
		}
	*/
	return c.String(http.StatusOK, "status"+c.Param("id"))
}

func matchesMove(c echo.Context) error {
	// to-do: verify if move is allowed (player turn, available spot,
	// unforseen rules (ko and super-ko))
	return c.String(http.StatusOK, "move"+c.Param("id"))
}
