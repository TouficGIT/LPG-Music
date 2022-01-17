package cmd

import (
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/print"
)

func JoinCommand(ctx framework.Context) {
	if ctx.Sessions.GetByGuild(ctx.Guild.ID) != nil {
		ctx.Reply(print.Join_AlreadyConnected)
		return
	}
	vc := ctx.GetVoiceChannel()

	if vc == nil {
		ctx.Reply(print.Join_NotInAChannel)
		return
	}

	sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
		Muted:    false,
		Deafened: true,
	})

	if err != nil {
		ctx.Reply(print.ReplyIssue)
		return
	}
	ctx.Reply("I've joined <#" + sess.ChannelId + ">")

}

/*

	var musicChannels string
	var channelMatch = make([]string, 0, 2)
	var connect bool

	channels, _ := ctx.Discord.GuildChannels(ctx.Guild.ID)

	for _, c := range channels {
		// Check if channel is a Music and a Vocal channel
		if strings.Contains(strings.ToLower(c.Name), "music") && c.Type == discordgo.ChannelTypeGuildVoice {
			channelMatch = append(channelMatch, c.ID)
			musicChannels += "\n\n<#" + c.ID + ">"
		}
	}

	if vc == nil {
		ctx.Reply("Tu dois être dans un channel Music pour utiliser cette commande!" + musicChannels)
		return
	}

	for _, c := range channelMatch {
		if c == vc.ID {
			connect = true
			break
		} else {
			connect = false
		}
	}

	if connect {
		sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
			Muted:    false,
			Deafened: true,
		})

		if err != nil {
			ctx.Reply("Oops ... il y eu un problème!")
			return
		}
		ctx.Reply("J'ai rejoint <#" + sess.ChannelId + ">")
	} else {
		ctx.Reply("Tu dois être dans un channel Music pour utiliser cette commande!" + musicChannels)
		return
	}

}
*/
