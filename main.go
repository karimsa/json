package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const debugEnabled = false

func debug(format string, args ...interface{}) {
	if debugEnabled {
		fmt.Fprintf(os.Stderr, "debug: "+format+"\n", args...)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func escape(val string) string {
	val = strings.Replace(val, "\n", "\\n", -1)
	val = strings.Replace(val, "\r", "\\r", -1)
	val = strings.Replace(val, "\t", "\\t", -1)
	val = strings.Replace(val, "$", "\\$", -1)

	return val
}

func boolAsStr(val bool) string {
	if val {
		return "true"
	}

	return "false"
}

func prefixed(prefix string, key string) string {
	if prefix == "" {
		return key
	}

	return prefix + "_" + key
}

func define(name string, value string) string {
	debug("define %s as \"%s\"", name, value)
	return name + "=\"" + value + "\" "
}

func exportMaybeArray(data interface{}, prefix string) (output string) {
	dat, ok := data.([]interface{})

	if !ok {
		return exportMaybeObject(data, prefix)
	}

	debug("parsing value at %s as array", prefix)
	output += define(prefix+"_length", strconv.Itoa(len(dat)))

	for index, value := range dat {
		key := prefixed(prefix, strconv.Itoa(index))
		val, ok := exportMaybePrimitive(value, key)

		if ok {
			output += val
		} else {
			output += exportMaybeArray(value, key)
		}
	}

	return
}

func exportMaybePrimitive(data interface{}, prefix string) (string, bool) {
	switch val := data.(type) {
	case float64:
		debug("parsing value at %s as float", prefix)
		if val == float64(int64(val)) {
			return define(prefix, strconv.FormatFloat(val, 'f', 0, 64)), true
		}
		return define(prefix, strconv.FormatFloat(val, 'f', 15, 64)), true

	case bool:
		debug("parsing value at %s as bool", prefix)
		return define(prefix, boolAsStr(val)), true

	case string:
		debug("parsing value at %s as string", prefix)
		return define(prefix, escape(val)), true

	default:
		return "", false
	}
}

func exportMaybeObject(data interface{}, prefix string) (output string) {
	dat, ok := data.(map[string]interface{})

	debug("parsing value at %s as object", prefix)

	if !ok {
		check(errors.New("Failed to convert key with prefix: " + prefix))
	}

	for key, value := range dat {
		pkey := prefixed(prefix, key)
		if val, ok := exportMaybePrimitive(value, pkey); ok {
			output += val
		} else {
			output += exportMaybeArray(value, pkey)
		}
	}

	return
}

func main() {
	rootPrefix := ""

	if len(os.Args) == 2 {
		rootPrefix = os.Args[1]
	}
	
	if strings.Contains(rootPrefix, "-")|| len(os.Args) > 2 {
		fmt.Println("usage: json [prefix] < file")
		fmt.Println("Or: cat file | json [prefix]")
		fmt.Println("Or whatever. Just pipe me stuff.")
		return
	}

	buffer, err := ioutil.ReadAll(os.Stdin)
	check(err)

	var data map[string]interface{}
	json.Unmarshal(buffer, &data)

	output := exportMaybeArray(data, rootPrefix)
	fmt.Println(output)
}
