package server

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"sandbox/internal/repository"
	"strconv"
)

const projectID = "advent"

func Root(c echo.Context) error {
	data := map[string]interface{}{}
	return c.Render(http.StatusOK, "index", data)
}

func FetchMember(c echo.Context) error {
	ctx := c.Request().Context()
	data := map[string]interface{}{}
	client, err := firestore.NewClient(ctx, projectID)

	members, err := repository.FetchAllMembers(ctx, client)
	if err != nil {
		return err
	}

	data["members"] = members
	return c.Render(http.StatusOK, "index", data)
}

func PutMember(c echo.Context) error {
	ctx := c.Request().Context()
	client, _ := firestore.NewClient(ctx, projectID)
	data := map[string]interface{}{}
	r := c.Request()
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		fmt.Println(err)
	}
	repository.PutMember(ctx, client, &repository.Member{
		Id:   r.FormValue("id"),
		Name: r.FormValue("name"),
		Age:  int32(age),
	})

	data["ok"] = true

	return c.Render(http.StatusOK, "index", data)
}

func DeleteAllMember(c echo.Context) error {
	ctx := c.Request().Context()
	client, _ := firestore.NewClient(ctx, projectID)
	repository.DeleteAllMember(ctx, client)
	return c.Render(http.StatusOK, "index", nil)
}

func DeleteMember(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()

	client, _ := firestore.NewClient(ctx, projectID)
	repository.DeleteMember(ctx, client, r.FormValue("id"))
	return c.Render(http.StatusOK, "index", nil)
}
