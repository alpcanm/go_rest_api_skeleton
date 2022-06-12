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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var date, tag, eq, gt string = "date", "tag", "$eq", "$gt"

func GetAllRaffles(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var isTagThere bool
	//! filtreli arama mı değil mi kontrolü.
	if c.QueryParam("tags") != "" {

		isTagThere = true
	}
	//!gt parametresi pagination için gerekli. greater ın kısaltması. Örn: Greater than 10 dateTime i 10 dan büyük olanları getirir.
	gtFitlers := gtFilters(c.QueryParam("gt"))
	//! istek içerisindeki filtreleri tagFilters değikenine atar.
	tagFilters := tagFilters(c.QueryParam("tags"))
	//! date e göre sıralama value 1 önce küçük olanı getir demek. Yani ilk önce yakın tarihli gelicek.
	sortValue := bson.D{{Key: date, Value: 1}}
	//! set limit3 en fazla 3 tane getir demek sayıyı büyütürsek daha fazla getirir.
	opts := options.Find().SetSort(sortValue).SetLimit(3)
	//! bütün filtreleri ve sıralama opsiyonlarını kullanıp db den veriyi çektiğimiz fonksiyon.
	fetchedData, err := rafflesCollection.Find(ctx, filterCheck(gtFitlers, tagFilters, isTagThere), opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//! hata kontrolü
			return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		}
	}
	//! gelen raffle ları atayalcağımız sonuçlar listesini oluşturduk.
	var results []models.RaffleModel
	//! for döngüsüyle gelen dataları sırayla results arrayinin içine atıyoruz.
	for fetchedData.Next(ctx) {
		var singleRaffle models.RaffleModel
		if err = fetchedData.Decode(&singleRaffle); err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{Message: "error", Body: &echo.Map{"data": err.Error()}})
		}
		results = append(results, singleRaffle)
	}
	//! status 200 dönüp bitiriyoruz.
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": results}})
}

//! istek filtreli mi değil mi kontrolü.
func filterCheck(gtFilter primitive.E, tagFilters primitive.E, isTagThere bool) bson.D {
	if isTagThere {

		return bson.D{gtFilter, tagFilters}
	}

	return bson.D{gtFilter}
}

//! greater parametresindeki veriyi alıp mongodb veritipinde bir filtreye dönüştürdüğmüz fonksiyon.
func gtFilters(gtParam string) primitive.E {
	queryRaffleDate, err := strconv.Atoi(gtParam)
	if err != nil {
		panic(err)
	}
	greaterThan := bson.D{{Key: gt, Value: queryRaffleDate}}
	return primitive.E{Key: date, Value: greaterThan}

}

//! gelen filtreleri listeleyip mongodb veritipine dönüştürdüğümüz fonksiyon.
func tagFilters(primary string) primitive.E {
	tagList := strings.Split(primary, ",")

	var equalList []bson.D
	for _, b := range tagList {

		equalList = append(equalList, bson.D{{Key: tag, Value: bson.D{{Key: eq, Value: b}}}})
	}

	return primitive.E{Key: "$or", Value: equalList}
}
