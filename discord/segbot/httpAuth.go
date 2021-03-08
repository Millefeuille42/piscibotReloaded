package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	agent := discordAgent{
		session: gBot,
		channel: os.Getenv("BOT_DEV_CHANNEL"),
	}

	stateMap, stateOk := params["state"]
	codeMap, codeOK := params["code"]
	if !stateOk || len(stateMap) <= 0 || !codeOK || len(codeMap) <= 0 {
		w = writeErrorToResponse(w, http.StatusBadRequest, http.ErrHijacked.Error())
		return
	}

	state := strings.Split(stateMap[0], "-")
	if len(state) <= 1 {
		w = writeErrorToResponse(w, http.StatusBadRequest, http.ErrHijacked.Error())
		return
	}

	userFile, err := userLoadFile(state[1], agent)
	switch {
	case err != nil:
		w = writeErrorToResponse(w, http.StatusNotFound, "Not Found")
	case userFile.State != state[0]:
		w = writeErrorToResponse(w, http.StatusForbidden, http.ErrHijacked.Error())
	case userFile.Verified:
		w = writeErrorToResponse(w, http.StatusGone, "Link expired")
	default:
		userFile.Verified = true
		err = userWriteFile(userFile, agent, userFile.UserID)
		if err != nil {
			w = writeErrorToResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte("Authentication successful, you can now close this window"))
	}
}

func authLinkCreator(agent discordAgent) (string, string) {
	uri := "https://api.intra.42.fr/oauth/authorize"

	secureCodeByte := make([]byte, 8)
	_, err := rand.Read(secureCodeByte)
	if err != nil {
		logErrorToChan(agent, err)
		return "", ""
	}

	secureCode := ""
	for _, val := range secureCodeByte {
		secureCodeInt := int(val)
		secureCode = fmt.Sprintf("%s%d", secureCode, secureCodeInt)
	}

	uri = uri + "?client_id=" + os.Getenv("APPUID") +
		"&redirect_uri=http%3A%2F%2F" + os.Getenv("APP_HOST") +
		"%3A" + os.Getenv("SEGBOT_PORT") + "%2Fauth" +
		"&response_type=code" + "&scope=public" +
		"&state=" + secureCode + "-" + agent.message.Author.ID
	return uri, secureCode
}
