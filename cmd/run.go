package main

import (
	tgwechat "github.com/kvneight/tg-wechat"
	"github.com/sirupsen/logrus"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logrus.Error(err)
		return
	}
	// load plugins for this session
	// faceplusplus.Register(session)
	// replier.Register(session)
	// switcher.Register(session)
	// gifer.Register(session)
	// cleaner.Register(session)
	// laosj.Register(session)
	// joker.Register(session)
	// revoker.Register(session)
	// forwarder.Register(session)
	// system.Register(session)
	// youdao.Register(session)

	// enable by type example
	tgwechat.Register(session)

	for {
		if err := session.LoginAndServe(false); err != nil {
			logrus.Error("session err: %s", err)
			// for i := 0; i < 3; i++ {
			// 	logrus.Info("trying re-login with cache")
			// 	if err := session.LoginAndServe(true); err != nil {
			// 		logrus.Error("re-login error, %s", err)
			// 	}
			// 	time.Sleep(3 * time.Second)
			// }
			// if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.WEB_MODE); err != nil {
			// 	logrus.Error("create new sesion failed, %s", err)
			// 	break
			// }
		} else {
			logrus.Info("closed by user")
			break
		}
	}
}
