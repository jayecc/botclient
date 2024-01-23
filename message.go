package botclient

type msg struct {
	MessageType  string  `json:"msgtype"`
	Text         Message `json:"text,omitempty"`
	Markdown     Message `json:"markdown,omitempty"`
	Image        Message `json:"image,omitempty"`
	News         Message `json:"news,omitempty"`
	File         Message `json:"file,omitempty"`
	TemplateCard Message `json:"template_card,omitempty"`
}

type Message interface {
	GetType() string
}

type TextMessage struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

func (text *TextMessage) GetType() string {
	return Text
}

type MarkdownMessage struct {
	Content string `json:"content"`
}

func (markdown *MarkdownMessage) GetType() string {
	return Markdown
}

type ImageMessage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

func (image *ImageMessage) GetType() string {
	return Image
}

type NewsMessage struct {
	Articles []NewsMessageArticle `json:"articles"`
}

type NewsMessageArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

func (news *NewsMessage) GetType() string {
	return News
}

func (news *NewsMessage) AddArticle(article ...NewsMessageArticle) {
	if news.Articles == nil {
		news.Articles = make([]NewsMessageArticle, 0)
	}
	news.Articles = append(news.Articles, article...)
}

type FileMessage struct {
	MediaId string `json:"media_id"`
}

func (file *FileMessage) GetType() string {
	return File
}
