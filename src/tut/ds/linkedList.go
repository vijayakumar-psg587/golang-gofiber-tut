package ds

import (
	"fmt"

	"github.com/vijayakumar-psg587/golang-fiber-tut/src/tut/customTypes"
)

func CreateLinkedList() []customTypes.Node {
	nodeArr := make([]customTypes.Node, 0, 10)
	n1 := new(customTypes.Node)
	n1.SetNode("5")
	nodeArr = append(nodeArr, *n1)

	// Adding second val
	nodeArr = *customTypes.AddNode(&nodeArr, "22")

	fmt.Printf("Node Arr initial %v\n", nodeArr)
	return nodeArr
}

func CreateLinkedListFromArray(arr []string) []customTypes.Node {
	nodeArr := make([]customTypes.Node, 0, len(arr)) // here we are setting intial length and capacity of slice which is what is required
	for idx, val := range arr {
		fmt.Println("val", idx, val)
		nodeArr = *customTypes.AddNode(&nodeArr, val)
	}
	fmt.Println("Node linked list created from array", nodeArr)
	return nodeArr
}

func InsertNodeToLinkedList(defaulArr []string, loc int64) (*[]customTypes.Node, error) {
	nodeArr := CreateLinkedListFromArray(defaulArr)

	return customTypes.InsertNode(&nodeArr, "4", loc)
}
