package controllers

import (
	config_ "go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = config_.GetCollection(config_.DB, "fghg")

func InsertOneUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, models.Response{})
}

func GetAUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, models.Response{})
}

func UpdateAUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, models.Response{})
}
