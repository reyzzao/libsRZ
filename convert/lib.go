package convert

import (
	"encoding/json"
	"fmt"
)

func ConvertObjectInJson(obj any) string {
	resJson, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}
	res := string(resJson)
	return res
}
