package cmd

import (
	"strconv"

	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func VolumeCommand(ctx framework.Context) {
	if len(ctx.Args) > 0 {
		vol, err := strconv.Atoi(ctx.Args[0])
		if err == nil {
			if vol <= 100 && vol >= 0 {
				framework.VOLUME = float64(vol) / 100
				ctx.Reply(print.ReplyNewVolume(vol))
				return
			} else {
				ctx.Reply(print.Volume_Command)
			}
		}
	} else {
		ctx.Reply(print.Volume_Command)
	}
	return
}
