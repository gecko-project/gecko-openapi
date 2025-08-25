package main

import (
	"encoding/json"
	"fmt"
)

func PrintResp(resp interface{}, err error) {
	if err != nil {
		fmt.Printf("err = %v", err)
	} else if resp != nil {
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Printf("MarshalIndent fail = %\n", err)
		} else {
			fmt.Printf("request ok = %s\n", string(data))
		}
	}
}
