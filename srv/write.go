package srv

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add(HeaderContentType, ContentTypeJSONUTF8)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
