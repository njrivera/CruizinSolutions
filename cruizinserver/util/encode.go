package util

import (
	"encoding/json"
	"net/http"
)

func JSONEncode(i interface{}, w http.ResponseWriter) {
	js, _ := json.MarshalIndent(i, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
