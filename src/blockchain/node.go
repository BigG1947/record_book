package blockchain

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Node struct {
	ip     string
	port   int
	length int
	hash   []byte
	status bool
}

func GetNodeList() []Node {
	return []Node{
		{
			ip:     "localhost",
			port:   8082,
			length: 0,
			status: false,
		},
		{
			ip:     "127.0.0.1",
			port:   8083,
			length: 0,
			status: false,
		},
	}
}

func CheckNodesLive(list []Node) (bool, error) {
	client := &http.Client{}
	var countActiveNode int
	for i := range list {
		stringURL := fmt.Sprintf("http://%s:%d/blockchain/checkActivity", list[i].ip, list[i].port)
		addrNode := fmt.Sprintf("%s:%d", list[i].ip, list[i].port)
		req, err := http.NewRequest("GET", stringURL, strings.NewReader(addrNode))
		if err != nil {
			return false, err
		}
		resp, err := client.Do(req)
		if err != nil {
			list[i].status = false
			log.Printf("Node: %s:%d don`t active!\n", list[i].ip, list[i].port)
			continue
		}
		if resp.StatusCode == http.StatusOK {
			list[i].status = true
			countActiveNode++
			log.Printf("Node: %s:%d have status OK!\n", list[i].ip, list[i].port)
		}
	}
	if countActiveNode < 2 || countActiveNode%2 != 0 {
		return false, nil
	}
	return true, nil
}
