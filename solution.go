package main

import "fmt"
import "github.com/kyokomi/emoji"

func GetMessage() string {
	emoji_msg := emoji.Sprint("Hello :world_map:!")
	fmt.Println(emoji_msg)
	return ""
}

func main() {
	GetMessage()
}
