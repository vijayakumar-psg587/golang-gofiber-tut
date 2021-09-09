package customTypes

import (
	"fmt"
	"log"
	"os"

	gjson "github.com/tidwall/gjson"
)

type Deck []string
type Card struct {
	cardType string
	cardVal  []string
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println("Printing range val", i, card)
	}
}

// func (c Card) PrintCards() {
// 	for i, card := range c {
// 		fmt.Println("Printing range val", i, card)
// 	}
// }

func NewDeck() *[]Card {

	res, err := createCardsFromFile()
	if err != nil {
		log.Fatalln("Error in reading json from File:", err.Error())
	}
	fmt.Println(res)
	return &res
}

func createCardsFromFile() ([]Card, error) {
	// First read card types
	file, err := os.Open("./src/tut/files/cardTypes.json")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// fmt.Printf("reading file %v %v \n", file, reflect.TypeOf(file))
	cardTypeBytes := make([]byte, 4096)
	cardTypeCount, err := file.Read(cardTypeBytes)
	if err != nil {
		return []Card{}, err
	}

	// cardTypeString := string(cardTypeBytes[:cardTypeCount])
	cardTypeStringGSON := gjson.GetMany(string(cardTypeBytes[:cardTypeCount]), "cardTypes")
	// fmt.Println("val of cardType:", cardTypeStringGSON)
	// return gjson.Get(string(cardTypeBytes[:count]), "cardTypes"), nil

	// Now read card values
	valFile, valErr := os.Open("./src/tut/files/cardValues.json")
	if valErr != nil {
		fmt.Println("Cannot Open value file", err)
		return []Card{}, err
	}

	cardValBytes := make([]byte, 4096)
	valCount, err := valFile.Read(cardValBytes)
	if err != nil {
		fmt.Println("Cannot read values frm val json file", err)
		return []Card{}, err
	}

	// cardValString := string(cardValBytes[:valCount])
	cardValStringGSON := gjson.GetMany(string(cardValBytes[:valCount]), "cardValues")

	// fmt.Println("val of cardValString:", cardValStringGSON)
	//cardSlice := make([]Card, 0) // creating a slice of cards
	//cardMap := make(map[string][]string, 0)
	// IMP!!!, This is how to iterate over array of json, by default gjson.Result for the below is a json type
	// hence normal iteration over forloop doesnt work, hence started using ForEach

	/*for _, cTypeIdx := range gjson.GetMany(cardTypeString, "cardTypes") {
		// cardSlice[cTypeIdx].cardType = cType
		// for cVal, cValIdx := range cardValString {

		// }
		fmt.Println("vv:", cTypeIdx.Array())
		cTypeIdx.ForEach(func(key, value gjson.Result) bool {
			println(value.String())
			return true
		})
	}*/
	fmt.Printf("type of %v %T", cardTypeStringGSON, cardTypeStringGSON)
	cardSlices := make([]Card, 0)
	card := new(Card)
	cardV := []string{}

	cardTypeStringGSON[0].ForEach(func(typeKey, typeVal gjson.Result) bool {
		card.cardType = typeVal.Str
		cardV = make([]string, 0)
		cardValStringGSON[0].ForEach(func(key, value gjson.Result) bool {
			cardV = append(cardV, value.Str)
			card.cardVal = cardV
			return true
		})
		cardSlices = append(cardSlices, *card)
		return true
	})
	fmt.Println("complete card:", cardSlices)

	return cardSlices, nil

}
