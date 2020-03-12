package conf

import (
	"fmt"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	file, _ := os.OpenFile("config.json", os.O_CREATE|os.O_WRONLY, 0666)
	c := `{
		"int":123,
		"string": "sssss"
	}`
	file.Write([]byte(c))
	file.Close()
	config, err := GetConfig("config.json")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(config)
	os.Remove("config.json")
}
