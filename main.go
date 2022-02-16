package main

import "fmt"
import "github.com/kyokomi/emoji"

func GetMessage() string {
	emoji_msg := emoji.Sprint("Hello :world_map:!")
	return emoji_msg
}

func main() {
	fmt.Println(GetMessage())
}
