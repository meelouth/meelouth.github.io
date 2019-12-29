package server

import (
	"encoding/json"
	"fmt"
	"github.com/kurenkoff/bstu_backend/file"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kurenkoff/bstu_backend/model"
	"github.com/kurenkoff/bstu_backend/repository"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Rep  *repository.SQL
}

func New(sql *repository.SQL) *Server {
	srv := &Server{
		Echo: echo.New(),
		Rep:  sql,
	}

	srv.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost", "http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	srv.InitRouters()
	return srv
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	log.Fatal(s.Echo.Start(port))
}

func (s *Server) InitRouters() {
	s.Echo.GET("/games/:id", s.GetGame)
	s.Echo.GET("/games", s.GetGames)
	s.Echo.POST("/games", s.AddGame)
	s.Echo.GET("/images/:filename", s.GetImage)
}

func (s *Server) GetGame(ctx echo.Context) error {
	id := ctx.Param("id")

	game, err := s.Rep.GetGame(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	raw, err := json.Marshal(game)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, string(raw))
}

func (s *Server) GetGames(ctx echo.Context) error {
	games, err := s.Rep.GetGames()
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	raw, err := json.Marshal(games)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, string(raw))
}

func (s *Server) AddGame(ctx echo.Context) error {
	body := ctx.Request().Body
	defer body.Close()

	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println("raw body :")

	var req model.AddRequest
	err = json.Unmarshal(raw, &req)
	if err != nil {
		return ctx.String(http.StatusBadGateway, err.Error())
	}
	//fmt.Printf("%+v\n", req)

	req.Game.Icon, err = file.SaveImage(req.Game.Icon)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	req.Game.Icon = os.Getenv("HOST_ADR")  + "images/" + req.Game.Icon
	gameID, err := s.Rep.InsertGame(req.Game)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	senderID, err := s.Rep.InsertSender(req.Sender)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	_, err = s.Rep.InsertSendGame(model.SentGame{
		GameID:   gameID,
		SenderID: senderID,
	})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.String(http.StatusOK, "")
}

func (s *Server) GetImage(ctx echo.Context) error {
	filename := ctx.Param("filename")

	//raw, err := file.GetFile(filename)
	//if err != nil {
	//	return ctx.String(http.StatusBadRequest, err.Error())
	//}

	return ctx.File(filename)
}