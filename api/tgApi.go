package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kvneight/tg-wechat/config"
	"github.com/kvneight/tg-wechat/data"
	"github.com/sirupsen/logrus"
)

var httpClient http.Client

func init() {
	httpClient = http.Client{
		Timeout: time.Second * 5,
	}
}

func SendMessage(chatId int, content string) {

	jsonString, err := json.Marshal(data.TgMessage{ChatId: chatId, Text: content})
	if err != nil {
		return
	}
	logrus.Info("post message to chat ", chatId, " content: ", string(jsonString))
	res, err := httpClient.Post("https://api.telegram.org/bot"+config.Config.TgToken+"/sendMessage", "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		logrus.Error("send tg message failed:", err)
		GetMe()
		return
	}
	defer res.Body.Close()
	buf := make([]byte, 1024)
	_, err = res.Body.Read(buf)
	if err != nil {
		logrus.Error("parse tg message failed:", err)
		return
	}
	logrus.Info("send tg message result:", string(buf))
}
func GetMe() (id string, err error) {
	r, err := httpClient.Get("https://api.telegram.org/bot" + config.Config.TgToken + "/sendMessage")
	if err != nil {
		logrus.Error(" get me faild：", err)
		return "", err
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(data, m)
	if err != nil {
		logrus.Error(" parse me faild：", err)
		return "", err
	}
	return m["chat_id"].(string), nil
}
