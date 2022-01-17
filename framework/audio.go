package framework

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os/exec"

	"Work.go/LPG-Bot/LPGMusic/print"
	"github.com/bwmarrin/discordgo"
	"layeh.com/gopus"
)

const (
	CHANNELS   int = 2
	FRAME_RATE int = 48000
	FRAME_SIZE int = 960
	MAX_BYTES  int = (FRAME_SIZE * 2) * 2
)

var (
	VOLUME float64 = 1.0
	PAUSE  bool    = false
)

/*
this shit is messy and i don't fully understand it yet credit to github.com/bwmarrin's voice example for the base code
*/

func (connection *Connection) sendPCM(voice *discordgo.VoiceConnection, pcm <-chan []int16) {
	connection.lock.Lock()
	if connection.sendpcm || pcm == nil {
		connection.lock.Unlock()
		return
	}
	connection.sendpcm = true
	print.DebugLog("[DEBUG] Send PCM true", "[SERVER]")
	connection.lock.Unlock()
	defer func() {
		connection.sendpcm = false
		print.DebugLog("[DEBUG] Send PCM false", "[SERVER]")
	}()
	encoder, err := gopus.NewEncoder(FRAME_RATE, CHANNELS, gopus.Audio)
	if err != nil {
		print.CheckError("[ERROR] NewEncoder error", "[SERVER]", err)
		return
	}
	for {
		receive, ok := <-pcm
		if !ok {
			print.DebugLog("[DEBUG] PCM channel closed", "[SERVER]")
			return
		}
		opus, err := encoder.Encode(receive, FRAME_SIZE, MAX_BYTES)
		if err != nil {
			print.CheckError("[ERROR] Encoding error", "[SERVER]", err)
			return
		}
		if !voice.Ready || voice.OpusSend == nil {
			fmt.Printf("[ERROR] Discordgo not ready for opus packets. %+v : %+v", voice.Ready, voice.OpusSend)
			return
		}
		voice.OpusSend <- opus
	}
}

func (connection *Connection) Play(ffmpeg *exec.Cmd) error {
	if connection.playing {
		return errors.New("[ERROR] Song already playing")
	}
	connection.stopRunning = false
	stdout, err := ffmpeg.StdoutPipe()
	if err != nil {
		return err
	}
	buffer := bufio.NewReaderSize(stdout, 262144)
	err = ffmpeg.Start()
	print.DebugLog("[DEBUG] Start Ffmpeg", "[SERVER]")
	if err != nil {
		return err
	}
	connection.playing = true
	print.DebugLog("[DEBUG] Start playing", "[SERVER]")
	defer func() {
		connection.playing = false
		print.DebugLog("[DEBUG] Stop playing", "[SERVER]")
	}()
	connection.voiceConnection.Speaking(true)
	defer connection.voiceConnection.Speaking(false)
	if connection.send == nil {
		connection.send = make(chan []int16, 2)
	}
	print.DebugLog("[DEBUG] Start SendPCM function", "[SERVER]")
	go connection.sendPCM(connection.voiceConnection, connection.send)
	for {
		if connection.stopRunning {
			ffmpeg.Process.Kill()
			break
		}
		if PAUSE == true {
			continue
		}
		audioBuffer := make([]int16, FRAME_SIZE*CHANNELS)
		err = binary.Read(buffer, binary.LittleEndian, &audioBuffer)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			print.CheckError("[ERROR] End Of File", "[SERVER]", err)
			return nil
		}
		if err != nil {
			return err
		}

		for i := range audioBuffer {
			audioBuffer[i] = int16(math.Floor(float64(audioBuffer[i]) * VOLUME))
		}

		connection.send <- audioBuffer
	}
	return nil
}

func (connection *Connection) Stop() {
	connection.stopRunning = true
	connection.playing = false
	PAUSE = false
}
