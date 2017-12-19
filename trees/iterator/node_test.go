package iterator

import "strings"
import "encoding/json"

const example = `{
  "value": 10,
  "left": {
    "value": 6,
    "left": {
      "value": 2
    },
    "right": {
      "value": 8
    }
  },
  "right":{
    "value": 20,
    "left": {
      "value": 18,
      "left": {
        "value": 17
      }
    }
  }
}`

func treeFromJSON(input string) *Node {
	var out Node
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.Decode(&out)
	return &out
}
