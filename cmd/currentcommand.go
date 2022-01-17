package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func CurrentCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	current := sess.Queue.Current()
	if current == nil {
		ctx.Reply(print.ReplyPlaylistEmpty)
		return
	}
	ctx.Reply(print.Current_CurrentlyPlaying + current.Title + "`.")
}
