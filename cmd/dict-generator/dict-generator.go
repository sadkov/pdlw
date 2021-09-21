package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const charsIncluded = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func main() {

	var (
		dictSize  = 5
		dict1     = make(map[int]string, dictSize)
		strLength = 15
		fileName  = "dict1.json"
		err       error
		fmap      []byte
	)

	rand.Seed(time.Now().UnixNano())

	for w := 1; w <= dictSize; w++ {
		strLength = rand.Intn(25) + 10
		dict1[w] = generateString(strLength)
	}

	for key, value := range dict1 {
		fmt.Printf("[%d]=%s\n", key, value)
	}

	fmap, err = json.Marshal(dict1)
	if err != nil {
		fmt.Println("Dictionary Marshaling error:", err)
		return
	}

	//fmt.Println("fmap=", fmap)
	if len(fmap) > 0 {
		if err := os.WriteFile(fileName, fmap, 0644); err != nil {
			fmt.Println("os.WriteFile error:", err)
			return
		}
	}
}

func generateString(l int) string {

	var letterRunes = []rune(charsIncluded)
    var strLen = len(letterRunes)

	dest := make([]rune, l)
	for idx := range dest {
		dest[idx] = letterRunes[rand.Intn(strLen)]
	}
	return string(dest)
}
