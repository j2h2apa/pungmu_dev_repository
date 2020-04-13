package example_test

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	TestToken               = "1057300654:AAFEuLNx-9wgeLq2XyorkmuKIPmKjVGV57s"
	ChatID                  = 960527111
	ReplyToMessageID        = 41
	ExistingPhotoFileID     = "AgACAgUAAxkDAAMqXo6vHsNcjE1jFitH920yvIcj6JgAAg-pMRt-BHlUpXmvGNSEJQlEZ8BqdAADAQADAgADbQADm2EAAhgE"
	ExistingDocumentFileID  = "BQACAgUAAxkBAAMuXo6-MF4gbOtF2K_oRYs76yKYQwgAAqUAA34EeVRN_iJ1P6CR_RgE"
	ExistingAudioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingVideoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	ExistingStickerFileID   = "CAACAgIAAxkBAAMyXo7GB5MeNHhQfcbAjP1GzSujlMgAAgIAA8A2TxMI9W5F-oSnWRgE"
)

func getBot(t *testing.T) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(TestToken)
	bot.Debug = true
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	return bot, err
}

func TestNewBotAPI(t *testing.T) {
	_, err := tgbotapi.NewBotAPI(TestToken)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageForward(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewForward(ChatID, ChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewPhotoUpload(ChatID, "images/daisy-5009533_640.jpg")
	msg.Caption = "Photo test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
	bot, _ := getBot(t)

	data, _ := ioutil.ReadFile("images/daisy-5009533_640.jpg")
	b := tgbotapi.FileBytes{Name: "daisy-5009533_640.jpg", Bytes: data}

	msg := tgbotapi.NewPhotoUpload(ChatID, b)
	msg.Caption = "filebytes photo test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileReader(t *testing.T) {
	bot, _ := getBot(t)

	f, _ := os.Open("images/daisy-5009533_640.jpg")
	reader := tgbotapi.FileReader{Name: "daisy-5009533_640.jpg", Reader: f, Size: -1}

	msg := tgbotapi.NewPhotoUpload(ChatID, reader)
	msg.Caption = "filereader photo test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewPhotoUpload(ChatID, "images/daisy-5009533_640.jpg")
	msg.ReplyToMessageID = ReplyToMessageID
	msg.Caption = "reply photo test"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewDocumentUpload(ChatID, "images/daisy-5009533_640.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewDocumentShare(ChatID, ExistingDocumentFileID)
	msg.Caption = "document share test"
	_, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewAudioUpload(ChatID, "tests/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewVoiceUpload(ChatID, "tests/voice.ogg")
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithContact(t *testing.T) {
	bot, _ := getBot(t)

	contact := tgbotapi.NewContact(ChatID, "5551234567", "Test")

	if _, err := bot.Send(contact); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithLocation(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(tgbotapi.NewLocation(ChatID, 40, 40))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithVenue(t *testing.T) {
	bot, _ := getBot(t)

	venue := tgbotapi.NewVenue(ChatID, "A Test Location", "123 Test Street", 40, 40)

	if _, err := bot.Send(venue); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewVideoUpload(ChatID, "tests/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideoNote(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewVideoNoteUpload(ChatID, 240, "tests/videonote.mp4")
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewStickerUpload(ChatID, "images/daisy-5009533_640.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewStickerShare(ChatID, ExistingStickerFileID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	bot, _ := getBot(t)

	msg := tgbotapi.NewStickerShare(ChatID, ExistingStickerFileID)
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	bot, _ := getBot(t)

	file := tgbotapi.FileConfig{FileID: ExistingPhotoFileID}

	_, err := bot.GetFile(file)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

// 상단 입력 중 메세지 보임
func TestSendChatConfig(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(tgbotapi.NewChatAction(ChatID, tgbotapi.ChatTyping))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendEditMessage(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(tgbotapi.NewMessage(ChatID, "Test editing."))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	// edit := tgbotapi.EditMessageTextConfig{
	// 	BaseEdit: tgbotapi.BaseEdit{
	// 		ChatID : ChatID,
	// 		MessageID : msg.MessageID,
	// 	},
	// 	Text : "Updated Text.",
	// }

	// _, err = bot.Send(edit)
	// if err != nil {
	// 	t.Error(err)
	// 	t.Fail()
	// }
}
