package main

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
	"github.com/nsf/jsondiff"
)

func matchSubsetOfJson(path string, body []byte, expectedjson string, print func(a ...interface{})) (err error) {
	var v interface{}
	var expected interface{}
	var actual string
	print(fmt.Sprintf("Path: %s Expected: %s Body: %s", path, expectedjson, string(body)))
	// re-encode expected response
	if err := json.Unmarshal([]byte(expectedjson), &expected); err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	query, err := gojq.Parse(path)
	if err != nil {
		return err
	}
	code, err := gojq.Compile(query)
	if err != nil {
		return err
	}

	iter := code.Run(v)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		act, err := json.Marshal(v)
		if err != nil {
			return err
		}
		actual = string(act)
	}

	opts := jsondiff.DefaultConsoleOptions()
	d, s := jsondiff.Compare([]byte(actual), []byte(expectedjson), &opts)
	switch d {
	case jsondiff.FullMatch:
		return nil
		break
	case jsondiff.SupersetMatch:
		return nil
		break
	case jsondiff.NoMatch:
		return fmt.Errorf("No match for s=[%s] got=[%s] expected=[%s] result=[%s]", s, string(body), expectedjson, actual)
		break
	default:
		return fmt.Errorf("Unsupported match type, %d  for s=[%s] got=[%s] expected=[%s] result=[%s]", d, s, string(body), expectedjson, actual)
	}
	return nil
}
