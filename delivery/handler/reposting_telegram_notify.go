package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kpango/glg"
)

var (
	Token     string
	ChatId    string
	ParseMode string = "html"
)

func GetTokenBotWithChatID(tokenBot, ChatID string) {
	Token = tokenBot
	ChatId = ChatID
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id":    ChatId,
		"text":       text,
		"parse_mode": ParseMode,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	_ = glg.Info("Notifikasi '%s' terkirim", text)
	_ = glg.Info("Response JSON: %s", string(body))

	// Return
	return true, nil
}

// func RepostingNumResult(numOfSuccess, numOfFailed, totalReposting int) (string, error) {
// 	// flag redefined [ERROR]
// 	var message string
// 	message = `
// 	Halo sobat ITSO, Reposting saldo apex berdasarkan rekening aktif sudah selesai:

// 		=============================
// 		<b>Reposting Success:</b> <i>` + strconv.Itoa(numOfSuccess) + ` Rekening</i>
// 		<b>Reposting Failed:</b> <i>` + strconv.Itoa(numOfFailed) + ` Rekening</i>
// 		<b>Reposting Total:</b> <i>` + strconv.Itoa(totalReposting) + ` Rekening</i>
// 		=============================

// 		Your Helper :)
// 		<b>CT Support Assistant</b>`

// 	return message, nil
// }

func RepostingNumResult(numOfSuccess, numOfFailed, totalReposting int) (string, error) {
	// flag redefined [ERROR]
	var message string
	message = `
	Halo sobat ITSO, Reposting saldo untuk seluruh lembaga sudah selesai:

		=============================
		<b>Reposting Success:</b> <i>` + strconv.Itoa(numOfSuccess) + ` Rekening</i>
		<b>Reposting Failed:</b> <i>` + strconv.Itoa(numOfFailed) + ` Rekening</i>
		<b>Reposting Total:</b> <i>` + strconv.Itoa(totalReposting) + ` Rekening</i>
		=============================
		
		Your Helper :)
		<b>CT Support Assistant</b>`

	return message, nil
}
