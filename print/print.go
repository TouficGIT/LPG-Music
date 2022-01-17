package print

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	cTime = time.Now().Format("01-02-2006 15:04:05")

	DEBUG = false

	HelpMsg = `
**MUSIC:**
` + "`$add <ytb link>`" + ` : to add a youtube sound (video or playlist)
` + "`$current`" + ` : to display the current music
` + "`$join`" + ` : to make the bot join your channel
` + "`$leave`" + ` : to make the bot leave your channel
` + "`$pause`" + ` : to pause the playlist
` + "`$play`" + ` : to start the playlist
` + "`$play <ytb link>`" + ` : to add a sound and play it
` + "`$queue`" + ` : to show the playlist
` + "`$resume`" + ` : to stop the pause
` + "`$stop`" + ` : to stop the playlist
` + "`$skip`" + ` : to skip to the next music
` + "`$volume <0 to 100>`" + ` : to change the volume
` + "`$youtube <search>`" + ` : to search a music on Ytb

**NEED HELP?:**
` + "`$help`" + `: to show this message ^^ `

	EmbedHelp = &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0xFF8142, // Orange
		Description: HelpMsg,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Hi I'm LBGBot! I answer to the following commands:",
	}

	ReplyIssue         = "Oops ... something wrong happened! Try $help"
	ReplyPlaylistEmpty = "Playlist is empty! Add a music with `$add`."
	ReplyNotConnected  = "I'm not connected to a channel!\nUse `$join`."

	// [ADD]
	Add_Command         = "Add command usage: `$add <song>`\nValid entries: `youtube url`, `youtube id` \n\n (Works with videos or youtube playlists)"
	Add_LoadingPlaylist = "Loading the playlist..."
	Add_ReadyToPlay     = "I finished to add the playlist.\nUse`$play` to start! " +
		"And to show the current playlist, use `$queue`."

	// [CLEAR]
	Clear_PlaylistAlreadyEmpty = "The playlist is already empty"
	Clear_PlaylistCleared      = "Playlist has been cleared"

	// [CURRENT]
	Current_CurrentlyPlaying = "Playing: `"

	// [JOIN]
	Join_AlreadyConnected = "Already connected! Use `$leave` to make me leave."
	Join_NotInAChannel    = "You need to be on a voice channel to use this command!"

	// [PAUSE]
	Pause               = "Pause. To start again, use: `$resume`."
	Pause_AlreadyPaused = "Music already paused."

	// [PICK]
	Pick_Command     = "Pick command usage: `$pick <result number>`"
	Pick_TooMuch     = "You cannot pick more than 5 musics!"
	Pick_NotSearched = "You haven't searched for musics yet!"

	// [QUEUE]
	Queue_StopPlaying = "Stopping."
	Queue_EndPlaylist = "End of the Playlist."

	// [RESUME]
	Resume_AlreadyPlaying = "Music already playing"

	// [SKIP]
	Skip_MusicSkipped = "Music skipped!"

	// [VOLUME]
	Volume_Command = "Volume command usage: `$volume <0 to 100>` "

	// [YOUTUBE]
	Ytb_Command   = "Youtube command usage: `$youtube <search query>`"
	Ytb_PickMusic = "\n\nTo pick a music, use `!pick <number>`."
)

func ReplyNewVolume(vol int) string {
	return "Volume is now at " + fmt.Sprint(vol) + "%"
}

func CurrentlyPlaying(Title string, Thumbnail string, VOLUME float64) (string, string) {
	return "Playing: `" + Title + "`.\n\n Volume: `" + fmt.Sprint(VOLUME*100) + "%` \nPause: `$pause`", Thumbnail
}

func CheckError(msg string, user string, err error) {
	if err != nil {
		fmt.Println("Time: " + cTime + " || " + user + " || " + msg + " || " + err.Error())
		return
	}
}

func InfoLog(msg string, user string) {
	fmt.Println("Time: " + cTime + " || " + user + " || " + msg)
	return
}

func SetDebug(option string) {
	if option == "on" {
		DEBUG = true
	} else if option == "off" {
		DEBUG = false
	}
}

func DebugLog(msg string, user string) {
	if DEBUG == true {
		fmt.Println("Time: " + cTime + " || " + user + " || " + msg)
		return
	} else {
		return
	}

}
