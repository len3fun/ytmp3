package youtube

import (
	"fmt"
	yt "github.com/knadh/go-get-youtube/youtube"
	"os/exec"
	"strings"
)

func DownloadAudio(url string) (string, error) {
	if IsShortUrl(url) {
		splitedUrl := strings.Split(url, "/")
		url = splitedUrl[len(splitedUrl)-1]
	}
	video, err := yt.Get(url)
	if err != nil {
		fmt.Println("Get video error:", err)
		return "", err
	}

	videoTitle := video.Title
	filename := fmt.Sprintf("music/%s.mp3", videoTitle)

	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", filename, url)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Run command error:", err)
		return filename, err
	}

	return filename, nil
}
