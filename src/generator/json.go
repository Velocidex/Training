package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MarshalIndent(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	err = json.Indent(buf, b, "", " ")
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func StringIndent(v interface{}) string {
	result, err := MarshalIndent(v)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func Dump(v interface{}) {
	fmt.Println(StringIndent(v))
}
