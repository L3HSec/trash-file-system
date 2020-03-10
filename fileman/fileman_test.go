package fileman

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFileMan(t *testing.T) {
	man := NewManager("./")
	buffer := []byte{0x41, 0x42, 0x43, 0x44}
	_, err := man.SaveFile("test1.txt", bytes.NewReader(buffer))
	if err != nil {
		t.Fatal("cannot write")
	}
	_, err = man.SaveFile("test2.txt", bytes.NewReader(buffer))
	if err != nil {
		t.Fatal("cannot write")
	}

	reader, err := man.GetFile("test2.txt")
	if err != nil {
		t.Fatal("cannot read")
	}
	readBuffer, err := ioutil.ReadAll(reader)
	fmt.Println(string(readBuffer))
	reader.Close()

	man.DeleteFile("test1.txt")
	man.DeleteFile("test2.txt")
}
