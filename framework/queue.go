package framework

import "Work.go/LPG-Bot/LPGMusic/print"

type SongQueue struct {
	list    []Song
	current *Song
	Running bool
}

func (queue SongQueue) Get() []Song {
	return queue.list
}

func (queue *SongQueue) Set(list []Song) {
	queue.list = list
}

func (queue *SongQueue) Add(song Song) {
	queue.list = append(queue.list, song)
}

func (queue SongQueue) HasNext() bool {
	return len(queue.list) > 0
}

func (queue *SongQueue) Next() Song {
	song := queue.list[0]
	queue.list = queue.list[1:]
	queue.current = &song
	return song
}

func (queue *SongQueue) Clear() {
	queue.list = make([]Song, 0)
	queue.Running = false
	PAUSE = false
	queue.current = nil
}

func (queue *SongQueue) Start(sess *Session, callback func(string, string)) {
	queue.Running = true
	PAUSE = false
	for queue.HasNext() && queue.Running {
		song := queue.Next()
		curTitle, curThumbnail := print.CurrentlyPlaying(song.Title, song.Thumbnail, VOLUME)
		callback(curTitle, curThumbnail)
		sess.Play(song)
	}
	if !queue.Running {
		callback(print.Queue_StopPlaying, "")
	} else {
		callback(print.Queue_EndPlaylist, "")
	}
}

func (queue *SongQueue) Current() *Song {
	return queue.current
}

func (queue *SongQueue) Pause() {
	queue.Running = false
}

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}
