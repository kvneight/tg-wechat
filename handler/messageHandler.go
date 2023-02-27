package handler

import (
	tgApi "github.com/kvneight/tg-wechat/api"
	"github.com/kvneight/tg-wechat/config"
	"github.com/songtianyi/wechat-go/wxweb"
)

// HandleWechatMessage Handle the message from webchat
func HandleWechatMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {
	//TODO forward this message to tg
	if msgRecv.MsgType == wxweb.MSG_TEXT {
		handlTextMessage(session, msgRecv)
	}
}

func handlTextMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {
	if msgRecv.ToUserName == "filehelper" {
		tgApi.SendMessage(config.Config.ChatId, msgRecv.Content)
	} else if msgRecv.ToUserName == session.Bot.UserName {
		tgApi.SendMessage(config.Config.ChatId, "收到微信消息："+msgRecv.Content)
	}
}

// HandleTgMessage handle the message from tg
func HandleTgMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {

}
