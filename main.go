package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vijayakumar-psg587/golang-fiber-tut/src/tut/customTypes"
	"github.com/vijayakumar-psg587/golang-fiber-tut/src/tut/ds"
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
	testut()
}

func testut() {
	ds.CreateLinkedList()

	stringArr := []string{"11", "22", "33", "44", "55", "66"}
	n, err := ds.InsertNodeToLinkedList(stringArr, 2)
	if err == nil {
		fmt.Println("Created linked list from array and insertdata:", n)
	} else {
		log.Fatal(err)
		os.Exit(1)
	}

	//fmt.Println("Create linkedList from array:", ds.CreateLinkedListFromArray(stringArr))
	// cards := make(customTypes.Deck, 0)
	// sampleCardSlice := []string{"TT", "FF"}
	cards := customTypes.NewDeck()
	fmt.Println("Cards retrieved:", *cards)
	// cards = append(cards, sampleCardSlice...)
	// cards.Print()
}
