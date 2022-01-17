package cmd

import (
	"fmt"

	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
)

func AddCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply(print.Add_Command)
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	msg := ctx.Reply(print.Add_LoadingPlaylist)
	for _, arg := range ctx.Args {
		t, inp, err := ctx.Youtube.Get(arg)

		if err != nil {
			ctx.Reply(print.ReplyIssue)
			print.CheckError("[ERROR] Error getting input", "[SERVER]", err)
			return
		}

		switch t {
		case framework.ERROR_TYPE:
			ctx.Reply(print.ReplyIssue)
			fmt.Println("[ERROR] Error type", t)
			return
		case framework.VIDEO_TYPE:
			{
				video, err := ctx.Youtube.Video(*inp)
				if err != nil {
					ctx.Reply(print.ReplyIssue)
					print.CheckError("[ERROR] Error getting video1", "[SERVER]", err)
					return
				}
				song := framework.NewSong(video.Media, video.Title, video.Thumbnail, arg)
				sess.Queue.Add(*song)
				newEmbed := &discordgo.MessageEmbed{
					Description: "Sound: `" + song.Title + "` has been added to the playlist." +
						"\nUse `$play` to start the playlist!\nAnd to check the playlist, use `$queue`.",
					Author: &discordgo.MessageEmbedAuthor{},
					Color:  0xc22e2e,
				}
				ctx.Discord.ChannelMessageEditEmbed(ctx.TextChannel.ID, msg.ID, newEmbed)
				break
			}
		case framework.PLAYLIST_TYPE:
			{
				videos, err := ctx.Youtube.Playlist(*inp)
				if err != nil {
					ctx.Reply(print.ReplyIssue)
					print.CheckError("[ERROR] Error getting playlist", "[SERVER]", err)
					return
				}
				for _, v := range *videos {
					id := v.Id
					_, i, err := ctx.Youtube.Get(id)
					if err != nil {
						ctx.Reply(print.ReplyIssue)
						print.CheckError("[ERROR] Error getting video2", "[SERVER]", err)
						continue
					}
					video, err := ctx.Youtube.Video(*i)
					if err != nil {
						ctx.Reply(print.ReplyIssue)
						print.CheckError("[ERROR] Error getting video3", "[SERVER]", err)
						return
					}
					song := framework.NewSong(video.Media, video.Title, video.Thumbnail, arg)
					sess.Queue.Add(*song)
				}
				ctx.Reply(print.Add_ReadyToPlay)
				break
			}
		}
	}
}
