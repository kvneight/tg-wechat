package config

type Options struct {
	TgToken string `json "tg_token"`
	ChatId  int    `json "chat_id"`
}

var Config Options
