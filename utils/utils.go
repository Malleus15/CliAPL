package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func PrintMap(elem map[string]interface{}) {
	for key, element := range elem {
		fmt.Println(key, ": ", element)
	}
}

// ParseJsonToMap This functions accepts a byte array containing a JSON
func ParseJsonToMap(jsonBuffer []byte) ( []map[string]interface{}, error) {

	// We create an empty array
	var parameters []map[string]interface{}

	// Unmarshal the json into it. this will use the struct tag
	err := json.Unmarshal(jsonBuffer, &parameters)
	if err != nil {
		return nil, err
	}

	// the array is now filled with users
	return parameters, nil

}

func PrintListJson(m []map[string]interface{}){
	for i, s := range m {
		fmt.Println("\n"+strconv.Itoa(i+1)+")")
		PrintMap(s)
	}
}
