package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func PlayCommand(ctx framework.Context) {

	lenArgs := len(ctx.Args)

	if lenArgs == 0 {
		Play(ctx)
	} else if lenArgs == 1 {
		AddCommand(ctx)
		Play(ctx)
	}
}

func Play(ctx framework.Context) {
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

	go queue.Start(sess, func(msg string, thumbnail string) {
		ctx.ReplyThumbail(msg, thumbnail)
	})
}
