package tgwechat

import (
	"github.com/kvneight/tg-wechat/handler"
	"github.com/sirupsen/logrus"
	"github.com/songtianyi/wechat-go/wxweb"
)

// Register
// registers to wxweb module
func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(handleMessage), "tg-txt-forwarder")
	session.HandlerRegister.Add(wxweb.MSG_IMG, wxweb.Handler(handleMessage), "tg-img-forwarder")
	if err := session.HandlerRegister.EnableByName("tg-txt-forwarder"); err != nil {
		logrus.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("tg-img-forwarder"); err != nil {
		logrus.Error(err)
	}
}
func handleMessage(session *wxweb.Session, msgRecv *wxweb.ReceivedMessage) {

	logrus.Info("received message:from=", msgRecv.FromUserName)
	logrus.Info("received message:to=", msgRecv.ToUserName)
	logrus.Info("received message:", msgRecv.Content)
	logrus.Info("url=", msgRecv.Url)
	logrus.Info("msgId=", msgRecv.MsgId)
	logrus.Info("type=", msgRecv.MsgType)
	logrus.Info("subtype=", msgRecv.SubType)
	who := session.Cm.GetContactByUserName(msgRecv.Who)
	if who == nil {
		logrus.Info("unknown user name", msgRecv.Who)
		return
	}
	logrus.Info("who=", who)

	if session.Bot.UserName != who.UserName {
		// message from others to me
		handler.HandleWechatMessage(session, msgRecv)
	} else {
		logrus.Info("message sent to", who.DisplayName)
	}
}
