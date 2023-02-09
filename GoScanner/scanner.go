package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ports := []int{80, 443, 22, 8080}

	for _, port := range ports {
		address := fmt.Sprintf("mcdonalds.com:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				fmt.Println(address, "is closed")
			} else {
				fmt.Println(err.Error())
			}
			continue
		}
		conn.Close()
		fmt.Println(address, "is open")
	}
}
