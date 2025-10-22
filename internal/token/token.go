package token

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

type SimpleTokenizerV1 struct {
	StrToInt    map[string]int
	IntToStr    map[int]string
	encodeRegex *regexp.Regexp
	decodeRegex *regexp.Regexp
}

func NewSimpleTokenizerV1(vocab map[string]int) *SimpleTokenizerV1 {
	intToStr := make(map[int]string)
	ere := regexp.MustCompile(`[^,._!"()'\"?\s]+|[,._!"()'\"?\s]|--`)
	dre := regexp.MustCompile(`\s+([,.:;?!"()\'])`)

	for token, i := range vocab {
		intToStr[i] = token
	}

	return &SimpleTokenizerV1{
		StrToInt:    vocab,
		IntToStr:    intToStr,
		encodeRegex: ere,
		decodeRegex: dre,
	}
}

func (t *SimpleTokenizerV1) Encode(text string) ([]int, error) {
	preprocessed := t.encodeRegex.FindAllString(text, -1)
	var ids []int
	for _, value := range preprocessed {
		trim := strings.TrimSpace(value)
		if trim == "" {
			continue
		}
		id, ok := t.StrToInt[trim]
		if !ok {
			ids = append(ids, t.StrToInt["<|unk|>"])
			continue
		}
		ids = append(ids, id)
	}
	ids = append(ids, t.StrToInt["<|endoftext|>"])
	return ids, nil
}

func (t *SimpleTokenizerV1) Decode(ids []int) (string, error) {
	var build strings.Builder
	for _, v := range ids {
		c, ok := t.IntToStr[v]
		if c == "<|endoftext|>" {
			break
		}
		if !ok {
			return "", fmt.Errorf("id not in vocab :'%d'", v)
		}
		build.WriteString(c)
	}
	text := build.String()
	return t.decodeRegex.ReplaceAllString(text, "$1"), nil
}

func getText(url, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func Token() map[string]int {
	url := "https://raw.githubusercontent.com/rasbt/LLMs-from-scratch/main/ch02/01_main-chapter-code/the-verdict.txt"
	name := "the-verdict.txt"
	err := getText(url, name)
	if err != nil {
		return nil
	}
	text, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}

	re := regexp.MustCompile(`\w+|[^\w\s]+`)
	all_words := re.FindAllString(string(text), -1)
	uniqTokens := make(map[string]struct{})
	for _, word := range all_words {
		uniqTokens[word] = struct{}{}
	}

	sortedToken := make([]string, 0, len(uniqTokens))
	for token := range uniqTokens {
		sortedToken = append(sortedToken, token)
	}
	sort.Strings(sortedToken)
	vocab := make(map[string]int)
	for i, token := range sortedToken {
		vocab[token] = i
	}
	count := len(sortedToken)
	vocab["<|unk|>"] = count
	vocab["<|endoftext|>"] = count + 1
	return vocab
}
