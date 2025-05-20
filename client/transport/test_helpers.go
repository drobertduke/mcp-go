package transport

import "encoding/json"

func mustMarshal(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	if string(b) == "null" { // Ensure we send "null" not an empty RawMessage if v was nil or marshalled to null
		return json.RawMessage("null")
	}
	return b
}
