package main

import "fmt"
import "github.com/kyokomi/emoji"

func main() {
	emoji_msg := emoji.Sprint("Hello :world_map:!")
	fmt.Println(emoji_msg)
}
