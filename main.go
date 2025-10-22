package main

import (
	"fmt"

	"github.com/SyedAsadK/llm-from-scratch-go/internal/token"
)

func main() {
	vocab := token.Token()
	tokenizer := token.NewSimpleTokenizerV1(vocab)
	text := `It's the last he painted, you know,
Mrs. Gisburn said with pardonable pride.`
	ids, _ := tokenizer.Encode(text)
	fmt.Println("After encoding the ids")
	fmt.Println(ids)
	fmt.Println("After decoding the ids")
	fmt.Println(tokenizer.Decode(ids))
	text = "Hello, do you like tea?"
	encoded, _ := tokenizer.Encode(text)
	fmt.Println(tokenizer.Decode(encoded))
}
