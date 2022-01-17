package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func ResumeCommand(ctx framework.Context) {
	if !framework.PAUSE {
		ctx.Reply(print.Resume_AlreadyPlaying)
	} else {
		framework.PAUSE = false
	}
	return
}
