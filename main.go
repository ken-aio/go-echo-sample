package main

import (
	"net/http"
	"strconv"

	echoSwagger "github.com/ken-aio/echo-swagger"
	_ "github.com/ken-aio/go-echo-sample/docs"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// User is user strcut
type User struct {
	ID      int    `json:"id"`
	GroupID int    `json:"group_id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
}

// HTTPError is error response
type HTTPError struct {
	Code string `json:"code"`
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1314
// @BasePath /api/v1
func main() {
	e := echo.New()

	initRouting(e)

	e.Logger.Fatal(e.Start(":1314"))
}

func initRouting(e *echo.Echo) {
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", hello)
	e.GET("/api/v1/groups/:group_id/users", getUsers)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

// getUsers is getting users.
// @Summary get users
// @Description get users in a group
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID"
// @Param gender query string false "Gender" Enum(man, woman)
// @Success 200 {array} main.User
// @Failure 500 {object} main.HTTPError
// @Router /groups/{group_id}/users [get]
func getUsers(c echo.Context) error {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		return errors.Wrapf(err, "errors when group id convert to int: %s", groupIDStr)
	}
	gender := c.QueryParam("gender")
	users := []*User{}
	if gender == "" || gender == "man" {
		users = append(users, &User{ID: 1, GroupID: groupID, Name: "Taro", Gender: "man"})
		users = append(users, &User{ID: 2, GroupID: groupID, Name: "Jiro", Gender: "man"})
	}
	if gender == "" || gender == "woman" {
		users = append(users, &User{ID: 3, GroupID: groupID, Name: "Hanako", Gender: "woman"})
		users = append(users, &User{ID: 4, GroupID: groupID, Name: "Yoshiko", Gender: "woman"})
	}
	return c.JSON(http.StatusOK, users)
}
