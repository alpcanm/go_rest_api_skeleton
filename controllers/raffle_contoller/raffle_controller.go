package rafcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection = config.GetCollection(config.DB, "raffles")
var validate = validator.New()
var date, tag, eq, gt string = "date", "tag", "$eq", "$gt"

func InsertARaffle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var raffle models.RaffleModel

	if err := c.Bind(&raffle); err != nil {
		//bir hata varsa döner

		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	if err := validate.Struct(raffle); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}

	raffle.IsExpired = false
	result, err := collection.InsertOne(ctx, raffle)
	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//sonuç döner
	return c.JSON(http.StatusCreated, models.Response{Body: &echo.Map{"data": result}})

}

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

	resultList, err := collection.Find(ctx, filterCheck(gtFitlers, tagFilters, isTagThere), opts)
	if err != nil {
		panic(err)
	}
	var results []*models.RaffleModel
	if err := resultList.All(ctx, &results); err != nil {
		panic(err)
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
