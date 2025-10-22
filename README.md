# goLM - Language Model Tokenizer in Go

A simple tokenizer implementation in Go for building Large Language Models (LLMs) from scratch.

## Overview

This project implements a basic tokenizer (`SimpleTokenizerV1`) that can encode text into token IDs and decode token IDs back into text. It's based on the foundational concepts for building language models.

## Features

- **Text Encoding**: Convert text strings into token IDs
- **Text Decoding**: Convert token IDs back into readable text
- **Vocabulary Building**: Automatically builds a vocabulary from training text
- **Special Tokens**: Supports special tokens like `<|unk|>` for unknown words and `<|endoftext|>` for text boundaries

## Installation

Make sure you have Go 1.25.1 or later installed, then clone and build the project:

```bash
git clone https://github.com/SyedAsadK/goLM.git
cd goLM
go build
```

## Usage

Run the example:

```bash
go run main.go
```

### Example Code

```go
package main

import (
    "fmt"
    "github.com/SyedAsadK/llm-from-scratch-go/internal/token"
)

func main() {
    // Build vocabulary
    vocab := token.Token()
    
    // Create tokenizer
    tokenizer := token.NewSimpleTokenizerV1(vocab)
    
    // Encode text
    text := "It's the last he painted, you know, Mrs. Gisburn said with pardonable pride."
    ids, _ := tokenizer.Encode(text)
    fmt.Println("Encoded IDs:", ids)
    
    // Decode back to text
    decoded, _ := tokenizer.Decode(ids)
    fmt.Println("Decoded text:", decoded)
}
```

## Project Structure

- `main.go` - Example usage of the tokenizer
- `internal/token/token.go` - Tokenizer implementation
- `the-verdict.txt` - Sample text data for vocabulary building

## License

This project is part of learning materials for building LLMs from scratch.
