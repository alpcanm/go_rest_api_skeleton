package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSomeSubscribersFromRecentRaffle(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	recentRaffle := getRecentRaffle()
	var resultRaffle struct { //! sadece recent raffle içerisindeki listeyi çeker diğer datalara ihtiyaç olmadığığı için
		SubscriberList []models.SubscriberModel `bson:"subscriber_list" json:"subscriber_list"`
	}
	//! kazananları mı yoksa çekiliş ekranında rastgele gelenleri mi gösterileceğinin kontrolü. Eğer kazananları istiyorsa bu değer true gelir.
	winnersChoose := checkWinnersChooss(c.QueryParam("winnersChoose"))
	//! kaç kişi istendiğini kontrol eder. Eğer kazananları çekiyorsak bu değer 3 gelir.
	generateNumber := checkNumber(winnersChoose, c)
	//! sorgu çalışır.
	err := rafflesWithSubscriberCollection.FindOne(ctx, bson.M{"_id": recentRaffle.RaffleId}).Decode(&resultRaffle)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNoContent, models.Response{Message: err.Error()})
		}

	}
	// ! gelen datanını SubscriberList fieldindaki veriler başka bir değişkene atanır.
	subscriberList := resultRaffle.SubscriberList
	//! rasrgele bir şekilde subscriber list içindekilerin sırasını değiştiren fonksiyon.
	//! Örn içine 5 yazıldı 5 i 4,1,5,3,2 diye bir liste yapıp döndürüyor.
	indexesOfChoices := randomGenerator(len(subscriberList), generateNumber)
	//! indexesOfChoices listesinin içine sırayla subscriberları attığımız fonksiyon.
	choosens := chooseSomeSubscribers(indexesOfChoices, subscriberList)
	//! eğer kazananları seçiyorsa  WithIndexSubscriberModel içerisine kazananları atayan fonksiyon.
	//! Burada indexli olarak atmamızın sebebi aynı verileri veritabanına da kaydettiği için. yoksa rastgele kayıt yapıp
	//! birinci ikinci üçüncün yeri karılabilir.
	winners := setWinnersIndexForRaffleScreen(choosens)
	if winnersChoose {
		//!Eğer winnerChoose true ise kazananları raffleCollections içerisindeki o raffle içerisine winnerModel olarak ekleyecek.
		//! eğer ki winner çekiliyorsa direkt gönderilir.
		setWinnersToRaffleCollections(choosens, recentRaffle.RaffleId) //! kazananları db ye kayıt eden fonksiyon.

		return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": winners}})
	}
	indexedSomeSubscribers := setSomeSubscriberIndexForRaffleScreen(choosens)
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": indexedSomeSubscribers}})

}
func checkNumber(winnerChoose bool, c echo.Context) int {
	//! kaç adet subscriber çekileceğini kontrol eder.
	if winnerChoose {
		return 3
	}
	yNumberValue, err := strconv.Atoi(c.QueryParam("number"))
	if err != nil {
		panic(err)
	}
	return yNumberValue
}
func checkWinnersChooss(param string) bool {
	//! winnersChoose parametresinin true false değerini kontrol eder.
	if param == "true" {
		return true
	}
	return false
}
func chooseSomeSubscribers(indexList []int, subscriberList []models.SubscriberModel) []models.SubscriberModel {
	//! SubscriberList verisiyle gelen listedeki indexListteki değerlere göre seçimleri yapan fonksiyon
	var resultList []models.SubscriberModel
	for _, num := range indexList {
		resultList = append(resultList, subscriberList[num])
	}
	return resultList
}

func randomGenerator(x int, y int) []int {
	//! 0 ile x arasında y adet birbirinden farklı sayıları liste içinde döndüren fonksisyon
	rand.Seed(time.Now().UnixNano())
	var generatedNumbers []int

	isNumberIn := func(list []int, sayi int) bool {
		for _, num := range list {
			if num == sayi {

				return true
			}
		}
		return false
	}

	i := 0
	for i < y {
		randomNumber := rand.Intn(x)
		if !isNumberIn(generatedNumbers, randomNumber) {
			i++
			generatedNumbers = append(generatedNumbers, randomNumber)
		}
	}
	return generatedNumbers
}

func setWinnersToRaffleCollections(winners []models.SubscriberModel, objectId primitive.ObjectID) bool {
	//! kazananları raffle koleksiyonuna ekler
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var winnerModel models.WinnersModel
	winnerModel.WinnersModelId = primitive.NewObjectID()
	winnerModel.First = winners[0]
	winnerModel.Second = winners[1]
	winnerModel.Third = winners[2]

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "winners", Value: winnerModel}}}}
	_, err := rafflesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	return true

}

func setSomeSubscriberIndexForRaffleScreen(choosens []models.SubscriberModel) []models.WithIndexSubscriberModel {
	//!raffleScreendeki rastgele indexleri atar. seyircinin önüne çıkacak olan isimlerin indexlerini belirler.
	var result []models.WithIndexSubscriberModel
	indexList := randomGenerator(25, len(choosens))
	for i, num := range indexList {
		indexedSubsrcriberModel := models.WithIndexSubscriberModel{
			SubscribeModelId: choosens[i].SubscribeModelId,
			SubscriberId:     choosens[i].SubscriberId,
			RaffleNickName:   choosens[i].RaffleNickName,
			SubscribeDate:    choosens[i].SubscribeDate,
			Index:            num,
		}
		result = append(result, indexedSubsrcriberModel)
	}
	return result
}
func setWinnersIndexForRaffleScreen(choosens []models.SubscriberModel) []models.WithIndexSubscriberModel {
	//!kazananların indexlerini yazar
	var result []models.WithIndexSubscriberModel

	for i, _ := range choosens {
		indexedSubsrcriberModel := models.WithIndexSubscriberModel{
			SubscribeModelId: choosens[i].SubscribeModelId,
			SubscriberId:     choosens[i].SubscriberId,
			RaffleNickName:   choosens[i].RaffleNickName,
			SubscribeDate:    choosens[i].SubscribeDate,
			Index:            i,
		}
		result = append(result, indexedSubsrcriberModel)
	}
	return result
}
