package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var date, tag, eq, gt string = "date", "tag", "$eq", "$gt"

func GetRaffles(c echo.Context) error {
	// TODO: Düzeltilmesi gereken yer
	var isTagThere bool
	if c.QueryParam("tags") != "" {

		isTagThere = true
	}
	gtFitlers := gtFilters(c.QueryParam("gt"))

	tagFilters := tagFilters(c.QueryParam("tags"))

	sortValue := bson.D{{Key: date, Value: 1}}
	opts := options.Find().SetSort(sortValue).SetLimit(3)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fetchedData, err := rafflesCollection.Find(ctx, filterCheck(gtFitlers, tagFilters, isTagThere), opts)
	if err != nil {
		panic(err)
	}
	var results []models.RaffleModel

	for fetchedData.Next(ctx) {
		var singleRaffle models.RaffleModel
		if err = fetchedData.Decode(&singleRaffle); err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{Message: "error", Body: &echo.Map{"data": err.Error()}})
		}
		results = append(results, singleRaffle)
	}

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": results}})
}
func filterCheck(gtFilter primitive.E, tagFilters primitive.E, isTagThere bool) bson.D {
	if isTagThere {

		return bson.D{gtFilter, tagFilters}
	}

	return bson.D{gtFilter}
}

func gtFilters(ltParam string) primitive.E {
	queryRaffleDate, err := strconv.Atoi(ltParam)
	if err != nil {
		panic(err)
	}
	lowerThan := bson.D{{Key: gt, Value: queryRaffleDate}}
	return primitive.E{Key: date, Value: lowerThan}

}

func tagFilters(primary string) primitive.E {
	tagList := strings.Split(primary, ",")

	var equalList []bson.D
	for _, b := range tagList {

		equalList = append(equalList, bson.D{{Key: tag, Value: bson.D{{Key: eq, Value: b}}}})
	}

	return primitive.E{Key: "$or", Value: equalList}
}
