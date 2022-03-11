package prodcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"
	"strconv"
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
	if c.QueryParam("tag1") != "" {
		isTagThere = true
	}
	gtFitlers := gtFilters(c.QueryParam("gt"))
	tagFilters := tagFilters(c.QueryParam("tag1"), c.QueryParam("tag2"))

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
		fmt.Println("filtreli")
		return bson.D{gtFilter, tagFilters}
	}
	fmt.Println("filtresiz")

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

func tagFilters(tag1, tag2 string) primitive.E {

	equalTo1 := bson.D{{Key: eq, Value: tag1}}
	equalTo2 := bson.D{{Key: eq, Value: tag2}}
	tagQuery1 := bson.D{{Key: tag, Value: equalTo1}}
	tagQuery2 := bson.D{{Key: tag, Value: equalTo2}}
	return primitive.E{Key: "$or", Value: [2]bson.D{tagQuery1, tagQuery2}}
}

func DenemeProd(c echo.Context) error {
	eq := collection.Database().CreateCollection(context.TODO(), "adasdas")
	return c.JSON(200, models.Response{Body: &echo.Map{"data": eq}})
}
