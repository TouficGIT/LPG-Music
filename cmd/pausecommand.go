package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func PauseCommand(ctx framework.Context) {
	if framework.PAUSE {
		ctx.Reply(print.Pause_AlreadyPaused)
	} else {
		framework.PAUSE = true
		ctx.Reply(print.Pause)
	}
	return
}

/*
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	queue := sess.Queue
	if !queue.HasNext() {
		ctx.Reply(print.ReplyPlaylistEmpty)
		return
	}
	queue.Pause()
	ctx.Reply(print.ReplyPause)
*/
