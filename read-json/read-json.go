package readjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	myError "read-json/errors"
)

func ReadJson[T any](path string, v *T) *T {
	content, err := ioutil.ReadFile(path)
	myError.ErrorHandler(err)

	err = json.Unmarshal(content, v)
	myError.ErrorHandler(err)

	return v
}

func PrintJsonInFormat(v any) {
	vByte, err := json.MarshalIndent(v, "", "  ")
	myError.ErrorHandler(err)

	fmt.Printf("%s\n", string(vByte))
}
