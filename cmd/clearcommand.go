package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func ClearCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	if !sess.Queue.HasNext() {
		ctx.Reply(print.Clear_PlaylistAlreadyEmpty)
		return
	}
	sess.Queue.Clear()
	ctx.Reply(print.Clear_PlaylistCleared)
}
