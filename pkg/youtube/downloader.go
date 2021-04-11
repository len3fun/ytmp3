package youtube

import (
	"fmt"
	yt "github.com/knadh/go-get-youtube/youtube"
	"os/exec"
)

func DownloadAudio(url string) (string, error) {
	video, err := yt.Get(url)
	if err != nil {
		fmt.Println("Get video error:", err)
		return "", err
	}

	videoTitle := video.Title
	// TODO: get user name
	filename := fmt.Sprintf("/home/levi/music/%s.mp3", videoTitle)

	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", filename, url)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Run command error:", err)
		return filename, err
	}

	return filename, nil
}
