package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var i []string

func main() {
	// JSON data
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia", 18, 25000.98]}`)

	// JSON decoding
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatalf("json decoding error: %s\n", err)
		return
	}

	// Read JSON decoded structure data
	if m, ok := f.(map[string]interface{}); ok {
		for k, v := range m {
			MatchType(k, v)
		}
	}
}

func IsString(v interface{}) bool {
	switch v.(type) {
	case string:
		return true
	}
	return false
}

func MatchType(k interface{}, v interface{}) {
	switch vv := v.(type) {
	case string:
		if IsString(k) {
			fmt.Printf("%s(%T) is string %s\n", k, k, vv)
		} else {
			i = append(i, fmt.Sprintf("%s(%T)", vv, vv))
		}
	case int:
		if IsString(k) {
			fmt.Printf("%s(%T) is integer %d\n", k, k, vv)
		} else {
			i = append(i, fmt.Sprintf("%d(%T) ", vv, vv))
		}
	case float64:
		if IsString(k) {
			fmt.Printf("%s(%T) is float64 %f\n", k, k, vv)
		} else {
			i = append(i, fmt.Sprintf("%f(%T) ", vv, vv))
		}
	case []interface{}:
		fmt.Printf("%s(%T) is interface ", k, k)
		for ik, iv := range vv {
			MatchType(ik, iv)
		}
		fmt.Printf("[%s]\n", strings.Join(i, " "))
	default:
		fmt.Printf("%s is of a type I don't know how to handle %T\n", k, vv)
	}
}
