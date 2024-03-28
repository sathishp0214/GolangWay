package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var input = []byte(`{
	"input":"welcome to the world of <<params.provider>>. You will be working on <<params.project>>.",
	"params":[
		{
			"name":"provider",
			"value":"kubernetes"
		},
		{
			"name":"project",
			"value":"project-1"
		}
	]
}`)

func main() {
	fmt.Println("output Parsed from given payload --------- ", ParseJson(input))

}

func ParseJson(input []byte) string {
	var payload Payload
	err := json.Unmarshal(input, &payload)
	if err != nil {
		panic(err)
	}

	fmt.Println(payload)

	for _, j := range payload.Params {
		ParmMap[j.Name] = j.Value
	}

	regex := regexp.MustCompile(`\<\<params\.(\w+)\>\>`)

	tmp := regex.FindAllStringSubmatch(payload.Input, -1)

	outputString := regex.ReplaceAllString(payload.Input, "%s")

	fmt.Println(tmp, outputString)

	outputString = fmt.Sprintf(outputString, ParmMap[tmp[0][1]], ParmMap[tmp[1][1]])

	return outputString

}

type Payload struct {
	Input  string   `json:"input"`
	Params []Params `json:"params"`
}

type Params struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var ParmMap = map[string]string{}

// output := payload.Input
// // finalOutput := ""

// for _, i := range tmp {
// 	a, b := ParmMap[i[1]]
// 	fmt.Println(a, b)
// 	output = fmt.Sprintf(output)
// 	parse := fmt.Sprintf(`\<\<params\.%s\>\>`, i[1])
// 	regex1 := regexp.MustCompile(parse)
// 	output1 := regex1.ReplaceAllString(output, a)
// 	fmt.Println("============", output1)

// }

// "provider": "kubernetes", "project": "project-1"
