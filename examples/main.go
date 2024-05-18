package main

import "github.com/axiston/axiston-go"

func main() {
	client, err := axiston.NewClient()
	if err != nil {
		return
	}

	if err := client.Health(); err != nil {
		return
	}
}
