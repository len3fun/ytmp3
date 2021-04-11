package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"ytmp3/pkg/youtube"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You entered /start command")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know that command :(")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// TODO: youtube link checker
	downloadStartedMessage := tgbotapi.NewMessage(message.Chat.ID, "Video downloading has been started")
	_, err := b.bot.Send(downloadStartedMessage)
	if err != nil {
		fmt.Println("Send message to user error:", err)
	}
	filePath, err := youtube.DownloadAudio(message.Text)
	if err != nil {
		// TODO: logging
		errorMessage := tgbotapi.NewMessage(message.Chat.ID, "Some error occurred while was downloading")
		_, err := b.bot.Send(errorMessage)
		if err != nil {
			fmt.Println("Send message to user error:", err)
		}
	}

	mp3 := tgbotapi.NewAudioUpload(message.Chat.ID, filePath)

	_, err = b.bot.Send(mp3)
	if err != nil {
		// TODO: logging
		fmt.Println("Send message to user error:", err)
	}

	err = os.Remove(filePath)
	if err != nil {
		// TODO: logging
		fmt.Println("Remove file error:", err)
	}

}
