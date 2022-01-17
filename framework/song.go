package framework

import (
	"os/exec"
	"strconv"
)

type Song struct {
	Media     string
	Title     string
	Thumbnail string
	Duration  *string
	Id        string
}

func (song Song) Ffmpeg() *exec.Cmd {
	return exec.Command("ffmpeg", "-i", song.Media, "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac",
		strconv.Itoa(CHANNELS), "pipe:1")
}

func NewSong(media, title, thumbnail, id string) *Song {
	song := new(Song)
	song.Media = media
	song.Title = title
	song.Thumbnail = thumbnail
	song.Id = id
	return song
}
