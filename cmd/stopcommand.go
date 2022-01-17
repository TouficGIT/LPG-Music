package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func StopCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	if sess.Queue.HasNext() {
		sess.Queue.Clear()
	}
	sess.Stop()
}
