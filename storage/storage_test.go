package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestStorage(t *testing.T) {
	manager := NewStorageManager("a.db", "./uploadDir")
	payload := "asdfsaljdfklasjdfklasjdkfl"
	buffer := []byte(payload)
	file, err := manager.SaveFile("test", "fuck comment", bytes.NewReader(buffer))
	if err != nil {
		t.Fatal("fuck")
	}

	reader, err := manager.GetFile(file.ID)
	if err != nil {
		t.Fatal("fuck")
	}
	content, _ := ioutil.ReadAll(reader)
	if string(content) != payload {
		t.Fatal("fuck")
	}
	fmt.Println(file)
}
