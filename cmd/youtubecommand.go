package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
)

const result_format = "\n`%d` %s - %s"

var ytSessions ytSearchSessions = make(ytSearchSessions)

type (
	ytSearchSessions map[string]ytSearchSession

	ytSearchSession struct {
		results []framework.YTSearchContent
	}
)

func ytSessionIdentifier(user *discordgo.User, channel *discordgo.Channel) string {
	return user.ID + channel.ID
}

func YoutubeCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		ctx.Reply(print.Ytb_Command)
		return
	}
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		ctx.Reply(print.ReplyNotConnected)
		return
	}
	query := strings.Join(ctx.Args, "+")
	results, err := ctx.Youtube.Search(query)
	if err != nil {
		ctx.Reply(print.ReplyIssue)
		print.CheckError("[ERROR] Error searching youtube", "[SERVER]", err)
		return
	}
	if len(results) == 0 {
		ctx.Reply("I haven't find anything for: `" + query + "`.")
		return
	}
	buffer := bytes.NewBufferString("**Results** for `" + query + "`:\n")
	for index, result := range results {
		buffer.WriteString(fmt.Sprintf(result_format, index+1, result.Snippet.Title, result.ChannelTitle))
	}
	buffer.WriteString(print.Ytb_PickMusic)
	ytSessions[ytSessionIdentifier(ctx.User, ctx.TextChannel)] = ytSearchSession{results}
	ctx.Reply(buffer.String())
}
