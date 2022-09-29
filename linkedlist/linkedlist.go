package linkedlist

import (
	"dataStructure/node"
	"fmt"
)

type LinkedList struct {
	headNode *node.Node
}

func (l *LinkedList) AddToHead(data int) {
	nnode := node.Node{}
	nnode.SetData(data)
	if l.headNode == nil {
		l.headNode = &nnode
	} else {
		p := l.headNode
		l.headNode = &nnode
		l.headNode.SetNextNode(p)
	}

}

func (l *LinkedList) InsertNode(data int) {
	node_tmp := node.Node{}
	node_tmp.SetData(data)
	if l.headNode == nil {
		l.headNode = &node_tmp
	} else {
		p := l.headNode
		for p.GetNextNode() != nil {
			p = p.GetNextNode()
		}
		p.SetNextNode(&node_tmp)
	}

}

func (l *LinkedList) Print() {
	p := l.headNode
	for p != nil {
		fmt.Printf("->%v", p.GetData())
		p = p.GetNextNode()
	}

}

func (l *LinkedList) GetValue() int {
	return l.headNode.GetData()

}

func (l *LinkedList) GetNextNode() *node.Node {
	return l.headNode.GetNextNode()
}
