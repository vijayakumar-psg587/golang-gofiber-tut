package main

import (
	"fmt"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/common/models"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/common/utils"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/tut/customTypes"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/tut/ds"
	"log"
	"strings"

	// "reflect"
	// "strconv"
	// "github.com/gofiber/fiber/v2"
	// "github.com/vijayakumar-psg587/golang-fiber-tut/src/models"
)

func main() {

	// configModel := models.GetServerConfig()
	// fmt.Printf("Getting val from config %v \n", configModel)
	// app := fiber.New()
	// port := 3000
	// err := app.Listen("0.0.0.0:" + strconv.Itoa(port))
	// if err != nil {
	// 	fmt.Printf("Error in starting fiber service %v %v \n", reflect.TypeOf(err), err)
	// }
	// testut()
	//common()
	commonConfigFromEnv() // THis is a much shortend version of the above call, where viper is used here to get config from env
	//flatten()

}

func flatten() {
	strSliceOfSlice := make([][]string, 0, 10)
	slice1 := make([]string, 0, 10)
	slice1 = append(slice1, "testSlice1")
	slice2 := make([]string, 0, 10)
	slice2 = append(slice2, "testSlice2")

	strSliceOfSlice = append(strSliceOfSlice, slice1)
	strSliceOfSlice = append(strSliceOfSlice, slice2)
	acc := make([]interface {}, 0)
	err := utils.FlattenStringStruts(&acc, strSliceOfSlice, true)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("accumlated val:", acc)
	}
}

func common() {
	//ctx := new(context.Context)
	res, err := models.GetConfig()
	op := new(strings.Builder)
	if err != nil {
		log.Fatal(string(err))
	} else {

		op.Write((res)[:len(res)])
		fmt.Println("value:",op)

	}
}

func commonConfigFromEnv() {
	bArr := models.GetConfigFromEnv()
	fmt.Println(string(bArr[:len(bArr)]))
}

func testut() {
	ds.CreateLinkedList()

	stringArr := []string{"11", "22", "33", "44", "55", "66"}
	n, err := ds.InsertNodeToLinkedList(stringArr, 2)
	if err == nil {
		fmt.Println("Created linked list from array and insertdata:", n)
	} else {
		log.Fatal(err)

	}

	//fmt.Println("Create linkedList from array:", ds.CreateLinkedListFromArray(stringArr))
	// cards := make(customTypes.Deck, 0)
	// sampleCardSlice := []string{"TT", "FF"}
	cards := customTypes.NewDeck()
	fmt.Println("Cards retrieved:", *cards)
	// cards = append(cards, sampleCardSlice...)
	// cards.Print()
}
