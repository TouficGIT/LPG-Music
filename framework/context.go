package framework

import (
	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate
	Args         []string

	// dependency injection?
	Conf       *Config
	CmdHandler *CommandHandler
	Sessions   *SessionManager
	Youtube    *Youtube
}

func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel,
	user *discordgo.User, message *discordgo.MessageCreate, conf *Config, cmdHandler *CommandHandler,
	sessions *SessionManager, youtube *Youtube) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.Conf = conf
	ctx.CmdHandler = cmdHandler
	ctx.Sessions = sessions
	ctx.Youtube = youtube
	return ctx
}

func (ctx Context) Reply(content string) *discordgo.Message {
	embed := &discordgo.MessageEmbed{
		Description: "**" + content + "**",
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xc22e2e,
	}

	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
	if err != nil {
		print.CheckError("[ERROR] Error whilst sending message", "[SERVER]", err)
		return nil
	}
	return msg
}

func (ctx Context) ReplyThumbail(content string, thumbnail string) *discordgo.Message {
	embedThumbnail := &discordgo.MessageEmbedImage{
		URL: thumbnail,
	}
	embed := &discordgo.MessageEmbed{
		Description: "**" + content + "**",
		Author:      &discordgo.MessageEmbedAuthor{},
		Image:       embedThumbnail,
		Color:       0xc22e2e,
	}

	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
	if err != nil {
		print.CheckError("[ERROR] Error whilst sending message", "[SERVER]", err)
		return nil
	}
	return msg
}

func (ctx Context) ReplyPlaylist(content string, thumbnail string) *discordgo.Message {
	embedThumbnail := &discordgo.MessageEmbedThumbnail{
		URL:    thumbnail,
		Width:  168,
		Height: 94,
	}

	embed := &discordgo.MessageEmbed{
		Title:       "LPGBot.Music: PlayList",
		Thumbnail:   embedThumbnail,
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xc22e2e,
		Description: content,
	}

	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
	if err != nil {
		print.CheckError("[ERROR] Error whilst sending message", "[SERVER]", err)
		return nil
	}
	return msg
}

func (ctx *Context) GetVoiceChannel() *discordgo.Channel {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel
	}
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID == ctx.User.ID {
			channel, _ := ctx.Discord.State.Channel(state.ChannelID)
			ctx.VoiceChannel = channel
			return channel
		}
	}
	return nil
}
