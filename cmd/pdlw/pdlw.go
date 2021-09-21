package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Digits and minus are removed from the list used for generating the dict.
const charList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZабцファ"

func main() {

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	options, err := parseCommandLine(os.Args)
	if err != nil {
		log.Fatalf("error parsing command line: %v", err)
	}

	if options.generate {
		fmt.Println("Error: generate argument is not implemented yet")
		os.Exit(1)
	}

	if options.parse {
		result, err := parseDict(*options)
		if err != nil {
			log.Fatalf("error parsing dictionary: %v", err)
		}
		fmt.Printf("Final result = %s with size = %d \n", result, len([]rune(result)))
	}

	os.Exit(0)
}

func printUsage() {
	usage := `

 Description:

   pdlw prints the longest substring in a dictionary, containing only specified symbols
   Expected dictionary is map[int]string, expected file format is JSON
   Expected type of symbols file is one line plaintext file

 Usage:
 
    Read and parse dictionary:

       pdlw parse [dict1.json] [--symbols "abcd"]
       pdlw parse [dict1.json] [--symbols-from-file symbols.txt]

   To be implemented:

       pdlw generate dict1.json [--entries 5] [--entries-size-range 15 25]

  --entries N  Specifies number of dictionary entries (min:1, max:100, default:5)

  --entries-size-range X Y
               Specifies minimum and maximum length of dictionary value in symbols 
               ( X<Y, min:1, max:50, default: X=15, Y=25 )

`
	fmt.Printf("%s", usage)
}

func parseCommandLine(args []string) (*cliOpts, error) {

	var options = new(cliOpts)
	options.other = new(pdlwOptions)
	if args[1] == "parse" {
		options.parse = true
	}
	if args[1] == "generate" {
		options.generate = true
	}

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "--symbols":
			i++
			if i >= len(args) {
				return nil, fmt.Errorf("missing strings of symbols after --symbols")
			}
			options.other.symbols = args[i]
		case "--symbols-from-file":
			i++
			if i >= len(args) {
				return nil, fmt.Errorf("missing file name after --symbols-from-file")
			}
			options.other.mapFileName = args[i]
		default:
			options.other.dictFileName = args[i]
		}
	}

	options.setDefaults()

	return options, nil
}

func parseDict(opts cliOpts) (string, error) {
	var (
		mapJSON      []byte
		dictFileName = opts.other.dictFileName
		mapFileName = opts.other.mapFileName
		dictToParse = make(map[int]interface{})
		//runeList        = []rune(charList)
		byteList      []byte
		charListLocal string
		indexUnicode  = make(map[rune]bool)
		err           error
		longest   string
	)

	mapJSON, err = os.ReadFile(dictFileName)
	if err != nil {
		return "", fmt.Errorf("error reading file \"%s\": %s", dictFileName, err)
	}

	if len(opts.other.symbols) > 0 {
		charListLocal = opts.other.symbols
	} else {
		byteList, err = os.ReadFile(mapFileName)
		if err != nil {
			return "", fmt.Errorf("error reading file \"%s\": %s", mapFileName, err)
		}

		charListLocal = fmt.Sprintf("%s", byteList)

	}

	if indexUnicode = charToMap(charListLocal); indexUnicode == nil {
		return "", fmt.Errorf("string to Map conversion error ")
	}

	// extract dictionary from file content
	err = json.Unmarshal(mapJSON, &dictToParse)
	if err != nil {
		return "", fmt.Errorf("cannot unmarshal file content: %s", err)
	}

	// analyze dictionary
	for key, val := range dictToParse {

		fmt.Println("key=", key, " value:", val)

		l, err := findLongest(fmt.Sprintf("%v", val), indexUnicode)
		if err != nil {
			fmt.Println("cant find longest string:", err)
			fmt.Println("That happens, let's check the next dictionary entry...")
		} else {
			fmt.Printf("entry %d : longest = %s, size=%d \n", key, l, len([]rune(l)))
			if len(longest) < len(l) {
				longest = l
			}
		}

	}
	return longest, nil
}

// charToMap converts string to a map[rune]bool.
// All values in the map equals true.
// In case of problems, returns nil
func charToMap(str string) map[rune]bool {
	var runeMap = make(map[rune]bool)

	if len(str) == 0 {
		return nil
	}

	for _, v := range []rune(str) {
		runeMap[v] = true
	}
	if len(runeMap) > 0 {
		return runeMap
	} else {
		return nil
	}
}

type pdlwOptions struct {
	dictFileName string
	mapFileName  string
	symbols      string
	dictEntries  int
	sizeMin      int
	sizeMax      int
}

type cliOpts struct {
	parse    bool
	generate bool
	other    *pdlwOptions
}

func (o *cliOpts) setDefaults() {
	fmt.Printf("others in set: %v \n", o.other)
	switch {
	case o.other.dictFileName == "":
		o.other.dictFileName = "dict1.json"
		fallthrough

	case o.other.mapFileName == "":
		o.other.mapFileName = "symbols.txt"
		fallthrough

	case o.other.dictEntries == 0:
		o.other.dictEntries = 5
		fallthrough

	case o.other.sizeMin == 0:
		o.other.sizeMin = 15
		fallthrough

	case o.other.sizeMax == 0:
		o.other.sizeMax = 25
	}
}

// findLongest searches for longest sequence of symbols provided as map in a string.
// Map is used in search, and supposed to be in a form m[rune("あ")]=true
// ( boolean value doesn't matter ).
// In case of success, returns the longest sequence as a string and error=nil
func findLongest(str string, cm map[rune]bool) (string, error) {

	cmLen := len(cm)
	strLen := len(str)
	if cmLen == 0 || strLen == 0 {
		return "", fmt.Errorf("search string or list of characters are empty")
	}

	var from, to, size = 0, 0, 0
	var start, finish, longest = 0, 0, 0
	var prevMatch = true

	for i, c := range str {
		var match bool
		if !prevMatch {
			from, to = i, i
		}
		if _, match = cm[c]; match {
			size++
			to = to + len(string(c))
			prevMatch = match
		} else {
			if size > 0 {
				if size >= longest {
					start, finish, longest = from, to, size
				}
				from, size = to, 0
			}
			prevMatch = match
		}
	}

	if size > 0 {
		if size >= longest {
			start, finish, longest = from, to, size
		}
	}

	if longest == 0 {
		return "", fmt.Errorf("none of string characters found in map")
	}
	//fmt.Println("Longest: ",str[start:finish], start, finish, longest)
	return str[start:finish], nil
}
