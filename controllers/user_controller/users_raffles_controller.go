package usercont

import (
	"context"
	rafcont "go_rest_api_skeleton/controllers/raffle_contoller"
	"go_rest_api_skeleton/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SelectUsersRaffles(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result models.UserModel
	// doğrulanan tokenın içerisindeki uid parametresini alır string e dönüştürür
	uid := c.Param("uid")

	gt, errInt := strconv.Atoi(c.QueryParam("gt"))
	if errInt != nil {
		panic(errInt)
	}
	// filtreye göre istenilen veriyi getirir ve result değişkenine atar.
	err := collection.FindOne(ctx, bson.D{{Key: "uid", Value: uid}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"error": err.Error()}})
		}
		panic(err)
	}

	var raffleIdList []primitive.ObjectID
	for _, value := range result.SubscribedRaffles {
		raffleIdList = append(raffleIdList, value.RaffleId)
	}

	var usersRaffelList models.UsersRaffleList

	usersRaffelList.RaffleList = rafcont.GetFilteredRaffles(gt, raffleIdList)
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": usersRaffelList}})
}
