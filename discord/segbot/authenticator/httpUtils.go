package authenticator

import (
	"encoding/json"
	"net/http"
)

// parseMessage Unmarshal the message
func parseMessage(data []byte) MessageList {
	ret := MessageList{}

	err := json.Unmarshal(data, &ret)
	if err != nil {
		return MessageList{}
	}
	return ret
}

// writeErrorToResponse Writes error + error code to http writer
func writeErrorToResponse(w http.ResponseWriter, errCode int, errMessage string) http.ResponseWriter {
	w.WriteHeader(errCode)
	_, _ = w.Write([]byte(errMessage))
	return w
}
