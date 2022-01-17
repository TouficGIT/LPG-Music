package cmd

import (
	"fmt"
	"strconv"

	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
)

const (
	invalid_song_format = "Invalid song number `%d`. Min: 1, max: %d"
)

func PickCommand(ctx framework.Context) {
	argsLen := len(ctx.Args)
	if argsLen == 0 {
		ctx.Reply(print.Pick_Command)
		return
	}
	if argsLen > 5 {
		ctx.Reply(print.Pick_TooMuch)
		return
	}
	identifier := ytSessionIdentifier(ctx.User, ctx.TextChannel)
	var ytSession ytSearchSession
	var ok bool
	if ytSession, ok = ytSessions[identifier]; !ok {
		ctx.Reply(print.Pick_NotSearched)
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	rLen := len(ytSession.results)
	var msg *discordgo.Message
	for i := 0; i < argsLen; i++ {
		num, err := strconv.Atoi(ctx.Args[i])
		if err != nil {
			ctx.Reply(print.ReplyIssue)
			print.CheckError("[ERROR] Error parsing int", "[SERVER]", err)
			return
		}
		if num < 1 || num > rLen {
			ctx.Reply(fmt.Sprintf(invalid_song_format, num, rLen))
			return
		}
		result := ytSession.results[num-1]
		ctx.Reply(print.Add_LoadingPlaylist)
		_, inp, err := ctx.Youtube.Get(result.ID.VideoID)
		video, err := ctx.Youtube.Video(*inp)
		song := framework.NewSong(video.Media, video.Title, video.Thumbnail, result.ID.VideoID)
		sess.Queue.Add(*song)
		if msg != nil {
			msg, err = ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, msg.Content+", `"+song.Title+"`")
		} else {
			msg = ctx.Reply("Added: `" + song.Title + "`")
		}
	}
	if !sess.Queue.Running {
		ctx.Reply(print.Add_ReadyToPlay)
	}
}

/*
ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, msg.ID, msg.Content+
			"Musique ajoutée à la playlist.\nUtilise **!play** pour commencer à jouer la playlist! Et pour voir la playlist, utilises **!queue**.")
*/
