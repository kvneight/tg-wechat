package handler

import "github.com/songtianyi/wechat-go/wxweb"

// HandleWechatMessage Handle the message from webchat
func HandleWechatMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {
	//TODO forward this message to tg
}

// HandleTgMessage handle the message from tg
func HandleTgMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {

}
