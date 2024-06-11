package botclient

import (
	"fmt"
	"gopkg.in/resty.v1"
	"io"
)

const (
	sendUrl        = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
	uploadMediaUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&msgtype=file"
	Text           = "text"
	Markdown       = "markdown"
	Image          = "image"
	News           = "news"
	File           = "file"
	All            = "@all"
)

type WxWorkBot struct {
	httpClient *resty.Client
	key        string
}

func New(key string) *WxWorkBot {
	return &WxWorkBot{
		httpClient: resty.New().SetHeader("User-Agent", "v1.0.0"),
		key:        key,
	}
}

func (bot *WxWorkBot) Send(message Message) error {

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	m := &msg{
		MessageType: message.GetType(),
	}

	switch message.GetType() {
	case Text:
		m.Text = message
	case Markdown:
		m.Markdown = message
	case Image:
		m.Image = message
	case News:
		m.News = message
	case File:
		m.File = message
	default:
		return fmt.Errorf("unsupported message type: %s", message.GetType())
	}

	request := bot.httpClient.NewRequest()
	request.SetHeader("Content-Type", "application/json")
	request.SetBody(m)
	request.SetResult(&result)
	response, err := request.Post(fmt.Sprintf(sendUrl, bot.key))
	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		return fmt.Errorf("send message failed, status code: %s", response.Status())
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("send message failed, err code: %d, err msg: %s", result.ErrCode, result.ErrMsg)
	}

	return nil
}

func (bot *WxWorkBot) UploadMedia(filepath string) (mediaId string, err error) {

	var result struct {
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt string `json:"created_at"`
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}

	request := bot.httpClient.NewRequest()
	request.SetFile("media", filepath)
	request.SetResult(&result)
	response, err := request.Post(fmt.Sprintf(uploadMediaUrl, bot.key))
	if err != nil {
		return "", err
	}

	if !response.IsSuccess() {
		return "", fmt.Errorf("upload media failed, status code: %s", response.Status())
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("upload media failed, err code: %d, err msg: %s", result.ErrCode, result.ErrMsg)
	}

	return result.MediaId, nil
}

func (bot *WxWorkBot) UploadMediaReader(fileName string, read io.Reader) (mediaId string, err error) {

	var result struct {
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt string `json:"created_at"`
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}

	request := bot.httpClient.NewRequest()
	request.SetFileReader("media", fileName, read)
	request.SetResult(&result)
	response, err := request.Post(fmt.Sprintf(uploadMediaUrl, bot.key))
	if err != nil {
		return "", err
	}

	if !response.IsSuccess() {
		return "", fmt.Errorf("upload media failed, status code: %s", response.Status())
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("upload media failed, err code: %d, err msg: %s", result.ErrCode, result.ErrMsg)
	}

	return result.MediaId, nil
}
