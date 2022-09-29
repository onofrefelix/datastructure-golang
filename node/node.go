package node

type Node struct {
	data     int
	nextNode *Node
}

func (node *Node) SetData(data int) {
	node.data = data
}

func (node *Node) SetNextNode(nextNode *Node) {
	node.nextNode = nextNode
}

func (node *Node) GetData() int {
	return node.data
}

func (node *Node) GetNextNode() *Node {
	return node.nextNode
}
