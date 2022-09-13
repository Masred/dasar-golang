package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func longJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	longJson("Masred")
	longJson(10)
	longJson(true)
	longJson([]string{"Masred", "Ganteng", "Banget"})
}
