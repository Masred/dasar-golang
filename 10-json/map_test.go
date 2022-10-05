package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","price":"222222"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func TestMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apple Mac Book Pro",
		"price": 222222,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
