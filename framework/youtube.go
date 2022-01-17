package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"Work.go/LPG-Bot/LPGMusic/print"
)

const (
	ERROR_TYPE    = -1
	VIDEO_TYPE    = 0
	PLAYLIST_TYPE = 1
)

type (
	videoResponse struct {
		Formats []struct {
			Url string `json:"url"`
		} `json:"formats"`
		Thumbnails []struct {
			Url string `json:"url"`
		} `json:"thumbnails"`
		Title string `json:"title"`
	}

	VideoResult struct {
		Media     string
		Title     string
		Thumbnail string
	}

	PlaylistVideo struct {
		Id string `json:"id"`
	}

	YTSearchContent struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
		ChannelTitle string `json:"channelTitle"`
	}

	ytApiResponse struct {
		Content []YTSearchContent `json:"items"`
	}

	Youtube struct {
		Conf *Config
	}
)

func (youtube Youtube) getType(input string) int {
	if strings.Contains(input, "upload_date") {
		return VIDEO_TYPE
	}
	if strings.Contains(input, "_type") {
		return PLAYLIST_TYPE
	}
	return ERROR_TYPE
}

func (youtube Youtube) Get(input string) (int, *string, error) {
	print.DebugLog("[DEBUG] Request:"+input+"", "[SERVER]")
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", input)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		print.CheckError("[ERROR] "+fmt.Sprint(err)+": "+stderr.String()+"", "[SERVER]", err)
		return ERROR_TYPE, nil, err
	}
	str := out.String()
	return youtube.getType(str), &str, nil
}

func (youtube Youtube) Video(input string) (*VideoResult, error) {
	var resp videoResponse
	err := json.Unmarshal([]byte(input), &resp)
	if err != nil {
		return nil, err
	}
	return &VideoResult{resp.Formats[0].Url, resp.Title, resp.Thumbnails[0].Url}, nil
}

func (youtube Youtube) Playlist(input string) (*[]PlaylistVideo, error) {
	lines := strings.Split(input, "\n")
	videos := make([]PlaylistVideo, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var video PlaylistVideo
		err := json.Unmarshal([]byte(line), &video)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return &videos, nil
}

/*func (youtube Youtube) OldGet(id string) (*VideoResult, error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting youtube info,", err)
		return nil, err
	}
    fmt.Println(string(out.Bytes()))
	var resp response
	json.Unmarshal(out.Bytes(), &resp)
	u := resp.Formats[0].Url
	return &VideoResult{u, resp.Title}, nil
}*/

func (youtube Youtube) Search(query string) ([]YTSearchContent, error) {
	resp, err := http.Get("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=20&q=" + query + "&type=video&key=" + "AIzaSyDa38B307tfk6sf0RzY4eY0cB3elgdVaWs")
	if err != nil {
		return nil, err
	}
	var apiResp ytApiResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)

	return apiResp.Content, nil
}

/*

ytApiResponse struct {
		Items []struct {

	}

func (youtube Youtube) buildUrl(query string) (*string, error) {
	base := "https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=20&q=" + query + "&type=video&key=" + "AIzaSyDa38B307tfk6sf0RzY4eY0cB3elgdVaWs"
	fmt.Println(base)
	address, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	address.RawQuery = params.Encode()
	str := address.String()
	return &str, nil
}

func (youtube Youtube) Search(query string) (*ytApiResponse, error) {
	addr, err := youtube.buildUrl(query)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(*addr)
	if err != nil {
		return nil, err
	}
	var apiResp ytApiResponse
	json.NewDecoder(resp.Body).Decode(&apiResp)
	return &apiResp, nil
}

*/
