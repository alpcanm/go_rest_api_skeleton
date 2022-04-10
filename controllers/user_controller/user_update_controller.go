package usercont

import (
	"context"
	"encoding/json"
	"go_rest_api_skeleton/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func UserUpdateController(c echo.Context) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"data": err.Error()}})
	}
	uid := c.Param("uid")

	filter := bson.D{{Key: "uid", Value: uid}}

	update := bson.M{"$set": bson.D{{Key: "name", Value: json_map["name"]}, {Key: "raffle_nick_name", Value: json_map["raffle_nick_name"]}}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": result}})
}
