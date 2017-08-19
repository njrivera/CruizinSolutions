package util

import (
	"encoding/json"
	"net/http"
)

func JSONEncode(i interface{}, w http.ResponseWriter) {
	if err, ok := i.(error); ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		js, _ := json.MarshalIndent(i, "", "  ")
		w.Write(js)
	}
}
