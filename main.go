package main

import (
	"log"
	"net/http"
	"strconv"

	echoSwagger "github.com/ken-aio/echo-swagger"
	_ "github.com/ken-aio/go-echo-sample/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

// User is user strcut
type User struct {
	ID      int    `json:"id"`
	GroupID int    `json:"group_id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
}

// Req is request body
type Req struct {
	Query string `json:"query"`
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

	initMiddleware(e)
	initRouting(e)

	e.Logger.Fatal(e.Start(":1314"))
}

func initMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header:my-header}` + "\n",
	}))
	//e.Use(myMiddleware)
}

func initRouting(e *echo.Echo) {
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", hello)
	e.POST("/", postHello)
	e.PUT("/", putHello)
	e.DELETE("/", deleteHello)
	e.GET("/param", getParam)
	e.POST("/param", postParam)
	e.GET("/api/v1/groups/:group_id/users", getUsers)
}

func hello(c echo.Context) error {
	log.Println("hello action")
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func postHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "post"})
}

func putHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "put"})
}

func deleteHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "delete"})
}

func getParam(c echo.Context) error {
	r := &Req{}
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "good get", "param": r.Query})
}

func postParam(c echo.Context) error {
	r := &Req{}
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "good post", "param": r.Query})
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

func myMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("before action")
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after action")
		return nil
	}
}
