package jsonpatch

import (
	"encoding/json"
	"io"
)

type OperationType string

const (
	ADD     OperationType = "ADD"
	REMOVE                = "REMOVE"
	REPLACE               = "REPLACE"
	COPY                  = "COPY"
	MOVE                  = "MOVE"
	TEST                  = "TEST"
)

type Operation struct {
	Type  OperationType `json:"op"`
	Path  string        `json:"path"`
	From  string        `json:"from"`
	Value interface{}   `json:"value"`
}

type Patch struct {
	Operations []Operation `json:"operations"`
	Version    int         `json:"version"`
}

func Decode(reader io.Reader) (*Patch, error) {
	result := &Patch{}
	return result, json.NewDecoder(reader).Decode(result)
}
