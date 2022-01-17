package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"Work.go/LPG-Bot/LPGMusic/cmd"
	"Work.go/LPG-Bot/LPGMusic/framework"
	"Work.go/LPG-Bot/LPGMusic/logs"
	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
)

var (
	// Current time : for logs
	cTime = time.Now().Format("01-02-2006 15:04:05")

	// var bot discord
	lpgMusic   *discordgo.Session
	lpgbotUser *discordgo.User
	conf       *framework.Config

	// var youtube
	youtube        *framework.Youtube
	sessionManager *framework.SessionManager
	CmdHandler     *framework.CommandHandler

	//var generic
	PREFIX string
	err    error
)

func main() {
	fmt.Println("")
	print.InfoLog("[START] LPG Music is starting ...", "[SERVER]")

	// -- CONFIG lpgSound --

	// Read the config file
	conf = framework.LoadConfig("./config.json")
	if conf == nil {
		print.CheckError("[ERROR] Could not read config file", "[SERVER]", err)
		return
	}

	// Creation of lpgBot
	lpgMusic, err = discordgo.New("Bot " + conf.Token)
	if err != nil {
		print.CheckError("[ERROR] Creation of lpgSound", "[SERVER]", err)
		return
	}

	// Get Bot information
	lpgbotUser, err = lpgMusic.User("@me")
	if err != nil {
		print.CheckError("[ERROR] Error obtaining bot information", "[SERVER]", err)
		return
	}

	// Create logs file if doesn't exist
	err = logs.CheckAndCreate()
	if err != nil {
		print.CheckError("[ERROR] Check or create Logs", "[SERVER]", err)
		return
	}

	//msg handler
	PREFIX = conf.BotPrefix
	CmdHandler = framework.NewCommandHandler()
	lpgMusic.AddHandler(ready)
	lpgMusic.AddHandler(messageHandler)

	// -- START lpgSound --

	// Identify is sent during initial handshake with the discord gateway.
	// It is now required to retrieve events from Discord servers
	lpgMusic.Identify.Token = conf.Token
	lpgMusic.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	err = lpgMusic.Open()
	if err != nil {
		print.CheckError("[ERROR] Opening connection", "[SERVER]", err)
		return
	}
	ytbCommands()
	sessionManager = framework.NewSessionManager()
	youtube = &framework.Youtube{Conf: conf}
	print.InfoLog("[CONNECTED] LPG Music is connected !", "[SERVER]")

	// Set LPG Bot playing at !help
	lpgMusic.UpdateGameStatus(0, "Try $help")

	// Start listening commands from discord
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// Function to handle every message for the bot
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// -- GET INFOS --

	//Open log file
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		print.CheckError("[ERROR] Opening log file", "[SERVER]", err)
		return
	}
	defer f.Close()

	// Get Channel Id where message has been post
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		print.CheckError("[ERROR] Get Channel Id", "[SERVER]", err)
		return
	}

	// Get the guild (server)
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		print.CheckError("[ERROR] Get Guild Id", "[SERVER]", err)
		return
	}

	// Get the user (server)
	user, err := s.User(m.Author.ID)
	if err != nil {
		print.CheckError("[ERROR] Get User Id", "[SERVER]", err)
		return
	}

	// Check if user is admin
	var isAdmin bool = false
	for _, role := range m.Member.Roles {
		if role == "325384210133024768" {
			isAdmin = true
		}
	}

	// If user is a bot, we don't need to read the message
	if user.ID == lpgbotUser.ID || user.Bot {
		return
	}

	// -- COMMANDS --

	// Read message from User, and check if command exists
	if strings.HasPrefix(m.Content, PREFIX) {

		// Get command from the content
		content := m.Content[len(PREFIX):]
		args := strings.Fields(content)
		command := strings.ToLower(args[0])
		ytbCommand, found := CmdHandler.Get(command)
		ctx := framework.NewContext(s, g, c, user, m, conf, CmdHandler, sessionManager, youtube)

		// If Prefix but no command, return
		if len(content) < 1 {
			print.InfoLog("[INFO] len(content) < 1"+content+"", "[SERVER]")
			return
		}

		// Write into the logs
		print.InfoLog("[INFO] Command: "+command+" From: "+m.Author.Username+"", "[SERVER]")
		_, err = f.WriteString("Time: " + cTime + " || Message: " + command + " || From: " + m.Author.Username + "\n")
		print.CheckError("[ERROR] During WriteString logs", "[SERVER]", err)

		// If command exists in ytbCommand, run it
		if found {
			ctx.Args = args[1:]
			c := *ytbCommand
			c(*ctx)

			// Else run the non ytb command
		} else {

			// -- [ADMIN] --

			if isAdmin == true {
				switch command {
				case "debug":
					print.SetDebug(args[1])
					_, _ = s.ChannelMessageSend(m.ChannelID, "[ADMIN] Debug mode is "+args[1])
					return
				default:
				}
			}

			// -- [NOT ADMIN] --

			switch command {
			case "help":
				print.CheckError("[ERROR] Could not create channel between bot and user: ", user.Username, err)
				_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🧙")
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, print.EmbedHelp)

			default:
				_, _ = s.ChannelMessageSend(m.ChannelID, "I don't know this command "+m.Author.Username+". Try `$help`")
				_ = s.MessageReactionAdd(m.ChannelID, m.ID, "🤔")
			}
		}
	}
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "Try $help")
}

func ytbCommands() {
	// ??? means I haven't dug in
	CmdHandler.Register("add", cmd.AddCommand, "Add a song to the queue !add <youtube-link>")
	CmdHandler.Register("clear", cmd.ClearCommand, "empty queue???")
	CmdHandler.Register("current", cmd.CurrentCommand, "Name current song???")
	CmdHandler.Register("join", cmd.JoinCommand, "Join a voice channel !join attic")
	CmdHandler.Register("leave", cmd.LeaveCommand, "Leaves current voice channel")
	CmdHandler.Register("pick", cmd.PickCommand, "???")
	CmdHandler.Register("pause", cmd.PauseCommand, "Pause whats in the queue")
	CmdHandler.Register("play", cmd.PlayCommand, "Plays whats in the queue")
	CmdHandler.Register("queue", cmd.QueueCommand, "Print queue???")
	CmdHandler.Register("resume", cmd.ResumeCommand, "Resume the music")
	CmdHandler.Register("skip", cmd.SkipCommand, "Skip")
	CmdHandler.Register("stop", cmd.StopCommand, "Stops the music")
	CmdHandler.Register("volume", cmd.VolumeCommand, "Volume")
	CmdHandler.Register("youtube", cmd.YoutubeCommand, "???")
}
