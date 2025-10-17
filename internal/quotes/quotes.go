package quotes

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func GetQuotes(filePath string) ([]string, error) {
	var quotes []string

	fileQuotes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(fileQuotes), &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetRandomQuote(quotes []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	randomIndex := r.Intn(len(quotes))

	randomItem := quotes[randomIndex]

	return randomItem
}
