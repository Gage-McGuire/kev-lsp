package rpc

import (
	"testing"
)

type EncodedMessage struct {
	Testing bool `json:"testing"`
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	encoded := Encode(EncodedMessage{Testing: true})
	if encoded != expected {
		t.Fatalf("\n\n===Expected===\n%s\n\n====Actual====\n%s\n\n", expected, encoded)
	}
}

func TestDecode(t *testing.T) {
	encoded := "Content-Length: 17\r\n\r\n{\"method\":\"test\"}"
	method, content, err := Decode([]byte(encoded))
	if err != nil {
		t.Fatalf("Error decoding message: %s\n\n", err)
	}
	if len(content) != 17 {
		t.Fatalf("\n\n===Expected===\n%d\n\n====Actual====\n%d\n\n", 15, len(content))
	}
	if method != "test" {
		t.Fatalf("\n\n===Expected===\n%s\n\n====Actual====\n%s\n\n", "test", method)
	}
}
