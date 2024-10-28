package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moneyforward/aml_platform/internal/api/openapi/v1"
)

func main() {
	e := echo.New()

	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// register handlers
	// API
	api := handler{}
	openapi.RegisterHandlers(e, api)

	// View
	e.GET("/", handleIndex)
	e.GET("/users", handleUserList)

	// 8080ポートで Echo サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}

// handler for
type handler struct{}

func (h handler) GetUserList(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		&openapi.GetUserListResponse{
			{
				Id:   1,
				Name: "aml_admin_user",
			},
			{
				Id:   2,
				Name: "aml_user",
			},
		},
	)
}

func handleIndex(c echo.Context) error {
	t, err := template.ParseFiles("internal/view/index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
		return err
	}
	if err := t.Execute(c.Response().Writer, struct {
		Title   string
		Message string
		Time    time.Time
	}{
		Title:   "AML基盤",
		Message: "マネロン商事　銭洗太郎！",
		Time:    time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
		return err
	}
	return nil
}

type User struct {
	Id   int64
	Name string
}

func handleUserList(c echo.Context) error {
	users := []User{
		{Id: 1, Name: "aml_admin_user"},
		{Id: 2, Name: "aml_user"},
	}

	tmpl, err := template.ParseFiles("internal/view/user_list.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error loading template")
	}

	return tmpl.Execute(c.Response().Writer, users)
}
