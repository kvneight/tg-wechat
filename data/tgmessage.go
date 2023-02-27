package data

type TgMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
