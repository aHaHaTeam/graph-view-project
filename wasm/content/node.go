package content

type Node struct {
	name string
	data []byte
}

func NewNode(name string, data []byte) *Node {
	return &Node{name: name, data: data}
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (n *Node) Data() []byte {
	return n.data
}

func (n *Node) SetData(data []byte) {
	n.data = data
}
