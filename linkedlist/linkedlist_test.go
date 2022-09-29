package linkedlist_test

import (
	l "dataStructure/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnValue_WhenInsertValue_ThenShouldReceiveValue(t *testing.T) {

	lklist := l.LinkedList{}
	lklist.AddToHead(1)
	lklist.AddToHead(3)
	assert.Equal(t, 3, lklist.GetValue(), "valor armazenado na lista")
	assert.Equal(t, 1, lklist.GetNextNode().GetData(), "valor armazenado na lista")
}

func TestGivenAnValue_WhenInsertValueEnd_ThenShouldReceiveValue(t *testing.T) {

	lklist := l.LinkedList{}
	lklist.InsertNode(1)
	lklist.InsertNode(2)
	lklist.InsertNode(3)
	lklist.AddToHead(4)
	lklist.Print()
	assert.Equal(t, 4, lklist.GetValue(), "valor armazenado na lista")
	assert.Equal(t, 1, lklist.GetNextNode().GetData(), "valor armazenado na lista")
}
