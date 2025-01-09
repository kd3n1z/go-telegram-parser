package telegramparser

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type TelegramParser struct {
	secret []byte
}

func CreateParser(token string) TelegramParser {
	sha := hmac.New(sha256.New, []byte("WebAppData"))

	sha.Write([]byte(token))

	return TelegramParser{secret: sha.Sum(nil)}
}

func (parser *TelegramParser) Parse(query string) (WebAppInitData, error) {
	result := WebAppInitData{}

	values, err := url.ParseQuery(query)

	if err != nil {
		return result, err
	}

	var hash string

	keys := make([]string, 0, len(values))

	for key := range values {
		if key == "hash" {
			hash = values.Get(key)
			continue
		}

		keys = append(keys, key)
	}

	sort.Strings(keys)

	var stringBuilder strings.Builder

	for index, key := range keys {
		stringBuilder.WriteString(key)
		stringBuilder.WriteString("=")
		stringBuilder.WriteString(values.Get(key))

		if index < len(keys)-1 {
			stringBuilder.WriteString("\n")
		}
	}

	sha := hmac.New(sha256.New, parser.secret)

	sha.Write([]byte(stringBuilder.String()))

	if hex.EncodeToString(sha.Sum(nil)) != hash {
		return result, errors.New("hash does not match")
	}

	result.Hash = hash
	result.Signature = values.Get("signature")
	result.AuthDate, err = strconv.ParseInt(values.Get("auth_date"), 10, 64)

	if err != nil {
		return result, err
	}

	result.QueryId = values.Get("query_id")
	result.ChatType = values.Get("chat_type")
	result.ChatInstance = values.Get("chat_instance")
	result.StartParam = values.Get("start_param")
	result.CanSendAfter, _ = strconv.ParseInt(values.Get("can_send_after"), 10, 64)

	_ = json.Unmarshal([]byte(values.Get("user")), &result.User)
	_ = json.Unmarshal([]byte(values.Get("receiver")), &result.Receiver)

	return result, nil
}
