package prodcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"
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

func InsertAProduct(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product models.ProductModel

	if err := c.Bind(&product); err != nil {
		//bir hata varsa döner

		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"error": err.Error()}})
	}
	if err := validate.Struct(product); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"error": err.Error()}})
	}
	//user in uid sine mongodb _uid yerleştirid
	product.ProductId = primitive.NewObjectID()
	product.IsExpired = false
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"error": err.Error()}})
	}
	//sonuç döner
	return c.JSON(http.StatusCreated, models.Response{Body: &echo.Map{"data": result}})

}

func GetProducts(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	lowerThan := bson.D{{Key: "$lt", Value: 100}}
	equalTo1 := bson.D{{Key: "$eq", Value: "ev_esyasi"}}
	equalTo2 := bson.D{{Key: "$eq", Value: "elektronik"}}
	tagQuery1 := bson.D{{Key: "tag", Value: equalTo1}}
	tagQuery2 := bson.D{{Key: "tag", Value: equalTo2}}
	filter := bson.D{{Key: "draw_date", Value: lowerThan}, {Key: "$or", Value: [2]bson.D{tagQuery1, tagQuery2}}}
	sortValue := bson.D{{Key: "draw_date", Value: 1}}
	opts := options.Find().SetSort(sortValue).SetLimit(10) //+1 artan -1 azalana göre sıralar
	defer cancel()
	var results []*models.ProductModel

	resultList, _ := collection.Find(ctx, filter, opts)

	if err := resultList.All(ctx, &results); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": results}})
}
