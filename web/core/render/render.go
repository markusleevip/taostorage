package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DefaultEncoding = "UTF-8"
)

// return JSON format
func RenderJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset="+DefaultEncoding)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	fmt.Fprint(w)
}
