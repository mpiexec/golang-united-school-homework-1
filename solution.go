package main

import "fmt"
import "github.com/kyokomi/emoji"

func GetMessage() string {
	fmt.Println(emoji.Sprint("Hello :world_map:!"))
	return ""
}

func main() {
	GetMessage()
}
