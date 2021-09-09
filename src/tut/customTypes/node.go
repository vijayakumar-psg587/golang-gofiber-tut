package customTypes

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Value interface{}
}

func (n *Node) SetNode(val interface{}) {
	n.Next = nil

	n.Value = val

}

func (n *Node) GetNode() {

}

func InsertNode(nodeArr *[]Node, valToInsert interface{}, locationToInsert int64) (*[]Node, error) {

	arrLen := len(*nodeArr)
	nodeArrVal := *nodeArr

	var prevNode *Node
	var currentNode *Node
	if locationToInsert <= int64(arrLen) {
		// means value can be inserted successfully
		if locationToInsert == int64(int64(arrLen)) {
			// means we have to insert at the end
			nodeArrVal = *AddNode(nodeArr, valToInsert)
		} else {
			// somewhere in the middle
			// allNextArr := nodeArrVal[:locationToInsert-1]
			// allPrevArr := nodeArrVal[locationToInsert:]
			currentNode = &nodeArrVal[locationToInsert]
			prevNode = &nodeArrVal[locationToInsert-1]
			n1 := new(Node)
			n1.Value = valToInsert
			n1.Next = currentNode

			prevNode.Next = n1

			nodeArrVal = append(nodeArrVal[:locationToInsert], append([]Node{*n1}, nodeArrVal[locationToInsert:]...)...)

		}

	} else {

		return nil, fmt.Errorf("Location to insert - %v - cannot be greate than length of linkedList - %v", locationToInsert, arrLen)
	}

	return &nodeArrVal, nil
}

func AddNode(nodeArr *[]Node, val interface{}) *[]Node {

	var prevNode *Node
	var currentNode Node

	arrLen := len(*nodeArr)
	nodeArrVal := *nodeArr
	if arrLen == 1 {
		prevNode = &(nodeArrVal)[arrLen-1]
		currentNode.SetNode(val)
		currentNode.Next = nil

		prevNode.Next = &currentNode

		nodeArrVal = append(*nodeArr, currentNode)
	} else if arrLen == 0 {
		prevNode = new(Node)
		prevNode.SetNode(val)
		nodeArrVal = append(*nodeArr, *prevNode)
	} else {
		prevNode = &nodeArrVal[arrLen-1]
		currentNode.SetNode(val)
		currentNode.Next = nil

		prevNode.Next = &currentNode
		nodeArrVal = append(nodeArrVal, currentNode)
	}

	return &nodeArrVal
}
