package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func Encode(message any) string {
	content, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func Decode(message []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(message, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("invalid message - separator not found")
	}

	// Content-Length: <number>
	contentLengthBytes := bytes.Split(header, []byte{':', ' '})[1]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, errors.New("invalid message - content length not found")
	}

	var baseMessage BaseMessage
	err = json.Unmarshal(content[:contentLength], &baseMessage)
	if err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

// Split is a custom scanner split function for the RPC protocol.
// Ignore atEOF boolean argument.
func Split(message []byte, _ bool) (advance int, token []byte, err error) {
	// If the message does not contain the separator,
	// return nil to indicate that more data is needed.
	header, content, found := bytes.Cut(message, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	// Content-Length: <number>
	contentLengthBytes := bytes.Split(header, []byte{':', ' '})[1]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	// If the content is shorter than the content length
	// return nil to indicate that more data is needed.
	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, message[:totalLength], nil
}

func WriteResponse(writer io.Writer, message any) {
	response := Encode(message)
	writer.Write([]byte(response))
}
