package blockchain

type Node struct {
	IP   string
	Port int
}

func InitNodeList() *[]Node {
	return &[]Node{}
}
