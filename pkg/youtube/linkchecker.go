package youtube

import "strings"

func IsShortUrl(link string) bool {
	const youtubeShortSubStr = "https://youtu.be/"
	return strings.Contains(link, youtubeShortSubStr)
}

func IsLongUrl(link string) bool {
	const youtubeSubStr = "https://www.youtube.com/"
	return strings.Contains(link, youtubeSubStr)
}

func IsYoutubeLink(link string) bool {
	isYoutubeLink := IsShortUrl(link) || IsLongUrl(link)
	return isYoutubeLink
}
