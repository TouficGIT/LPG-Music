package cmd

import (
	"bytes"
	"fmt"
	"strconv"

	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

const (
	current_format  = "**Current music**: \n%s\n"
	song_format     = "\n`%03d` %s"
	next_format     = "\n**Next music**:"
	invalid_page    = "Invalid page `%d`. Min: `1`, max: `%d`"
	response_footer = "\n\nPage **%d** sur **%d**.\nTo go to the next page: `$queue <num page>`."
)

var Thumbnail string

func QueueCommand(ctx framework.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	queue := sess.Queue

	q := queue.Get()
	if len(q) == 0 && queue.Current() == nil {
		ctx.Reply(print.ReplyPlaylistEmpty)
		return
	}
	buff := bytes.Buffer{}
	if queue.Current() != nil {
		Thumbnail = queue.Current().Thumbnail
		buff.WriteString(fmt.Sprintf(current_format, queue.Current().Title))
	}
	queueLength := len(q)
	if len(ctx.Args) == 0 {
		var resp string
		buff.WriteString(fmt.Sprintf(next_format))
		if queueLength > 20 {
			resp = display(q[:20], buff, 1, 2, 0)
		} else {
			resp = display(q[:queueLength], buff, 1, 1, 0)
		}
		ctx.ReplyPlaylist(resp, Thumbnail)
		return
	}

	page, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		ctx.Reply("Invalid page `" + ctx.Args[0] + "`. Use: `$queue <page>`")
		return
	}
	pages := queueLength / 20
	if page < 1 || page > (pages+1) {
		ctx.Reply(fmt.Sprintf(invalid_page, page, pages+1))
		return
	}
	var lowerBound int
	if page == 1 {
		lowerBound = 0
	} else {
		lowerBound = (page - 1) * 20
	}
	upperBound := page * 20
	if upperBound > queueLength {
		upperBound = queueLength
	}
	slice := q[lowerBound:upperBound]
	ctx.ReplyPlaylist(display(slice, buff, page, pages, lowerBound), Thumbnail)
}

func display(queue []framework.Song, buff bytes.Buffer, page, pages, start int) string {
	for index, song := range queue {
		buff.WriteString(fmt.Sprintf(song_format, start+index+1, song.Title))
	}
	buff.WriteString(fmt.Sprintf(response_footer, page, pages))
	return buff.String()
}
