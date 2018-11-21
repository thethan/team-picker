package team_picker

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/jsonapi"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/thethan/team_picker/core"
	"net/http"
	"time"
)

const (
	contentJsonAPI    = "application/vnd.api+json"
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
)

type API struct {
	*core.Dependencies
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (api *API) SetupTokenRoutes(e *echo.Echo) {

	e.GET("/token", api.GetToken)

}

func (api *API) SetupRoutes(group *echo.Group) {
	group.Use(JsonAPIAccept)
	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(api.Dependencies.Config.JWTKeyString),
	}))

	group.GET("/teams", api.GetTeams)
}


func (api *API) GetToken(c echo.Context) error {
	api.Logger.Print("Token was hit")
	claims := &jwtCustomClaims{
		"Ethan",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(api.Dependencies.Config.JWTKeyString))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (api *API) GetTeams(c echo.Context) error {
	teams, err := GetTeamsFromYaml(api.Dependencies)
	if err != nil {
		api.Logger.Printf("Error in retrieving teams from yaml : %+v", err)
		return c.JSON(http.StatusExpectationFailed, "Error retrieving Teams")
	}
	return c.JSON(http.StatusOK, teams.Teams)
}

func JsonAPIAccept(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiAuthToken := c.Request().Header.Get(headerAccept)
		fmt.Printf("AuthToken %+v \n", apiAuthToken)
		if apiAuthToken == jsonapi.MediaType {
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "unauthorized"})
	}
}

func JsonAPIContentType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiAuthToken := c.Request().Header.Get(headerAccept)
		if apiAuthToken == jsonapi.MediaType {
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "unauthorized"})
	}
}
