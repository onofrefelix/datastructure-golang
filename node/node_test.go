package node_test

import (
	"dataStructure/node"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnValue_WhenInsertValue_ThenShouldReceiveValue(t *testing.T) {
	n1 := node.Node{}
	n1.SetData(1)
	n2 := node.Node{}
	n2.SetData(2)
	n2.SetNextNode(&n1)
	assert.Equal(t, 1, n2.GetNextNode().GetData(), "valor armazenado na lista")
	assert.Equal(t, 2, n2.GetData(), "valor armazenado na lista")

}
