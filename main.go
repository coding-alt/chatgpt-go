package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	transport, err := SetupProxy(config.ProxyOption)
	if err != nil {
		fmt.Println("Error setting up proxy:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入问题，或输入 'quit' 退出：")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if question == "quit" {
			break
		}

		_, err := CallAPI(question, config, transport)
		if err != nil {
			fmt.Println("Error calling API:", err)
			continue
		}

		fmt.Println()
	}
}
