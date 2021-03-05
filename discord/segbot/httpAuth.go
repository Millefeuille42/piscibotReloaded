package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func authLinkCreator(id string) (string, string) {
	uri := "https://api.intra.42.fr/oauth/authorize"

	agent := discordAgent{
		session: gBot,
		channel: os.Getenv("BOT_DEV_CHANNEL"),
	}

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
		"&response_type=code" + "&scope=public" + "&state=" + secureCode + "-" + id
	return uri, secureCode
}
