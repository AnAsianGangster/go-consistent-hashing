package main

import (
	"go-consistent-hashing/nodeStatus"

	"fmt"
)

func main() {
	nodeStatus.GetOneNodeStatus("5000")
	nodeStatus.ChangeNodeStatus("5000")
	fmt.Println("this is the main.go for go-consistent-hashing")
}
