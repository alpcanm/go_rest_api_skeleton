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

var collection *mongo.Collection = config.GetCollection(config.DB, "products")
var validate = validator.New()
var draw_date, tag, eq, gt string = "draw_date", "tag", "$eq", "$gt"

func InsertAProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product models.ProductModel

	if err := c.Bind(&product); err != nil {
		//bir hata varsa döner

		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	if err := validate.Struct(product); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//user in uid sine mongodb _uid yerleştirid
	product.ProductId = primitive.NewObjectID()
	product.IsExpired = false
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//sonuç döner
	return c.JSON(http.StatusCreated, models.Response{Body: &echo.Map{"data": result}})

}

func GetProducts(c echo.Context) error {
	// TODO: Düzeltilmesi gereken yer
	var isTagThere bool
	if c.QueryParam("tag1") != "" {
		isTagThere = true
	}
	gtFitlers := gtFilters(c.QueryParam("gt"))
	tagFilters := tagFilters(c.QueryParam("tag1"), c.QueryParam("tag2"))

	sortValue := bson.D{{Key: draw_date, Value: 1}}
	opts := options.Find().SetSort(sortValue).SetLimit(10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultList, err := collection.Find(ctx, filterCheck(gtFitlers, tagFilters, isTagThere), opts)
	if err != nil {
		panic(err)
	}
	var results []*models.ProductModel
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
	querDrawDate, err := strconv.Atoi(ltParam)
	if err != nil {
		panic(err)
	}
	lowerThan := bson.D{{Key: gt, Value: querDrawDate}}
	return primitive.E{Key: draw_date, Value: lowerThan}

}

func tagFilters(tag1, tag2 string) primitive.E {

	equalTo1 := bson.D{{Key: eq, Value: tag1}}
	equalTo2 := bson.D{{Key: eq, Value: tag2}}
	tagQuery1 := bson.D{{Key: tag, Value: equalTo1}}
	tagQuery2 := bson.D{{Key: tag, Value: equalTo2}}
	return primitive.E{Key: "$or", Value: [2]bson.D{tagQuery1, tagQuery2}}
}
