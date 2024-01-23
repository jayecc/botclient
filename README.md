# botclient
wx work bot client

```go
func TestText(t *testing.T) {
	msg := &TextMessage{
		Content:       "hello",
		MentionedList: []string{All},
	}
	bot := New("98fcbeba-ce84-451b-8b6d-ae162d769fd6")
	if err := bot.Send(msg); err != nil {
		t.Error(err)
	}
}

func TestUploadMedia(t *testing.T) {
	bot := New("98fcbeba-ce84-451b-8b6d-ae162d769fd6")
	mid, err := bot.UploadMedia("go.mod")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(mid)
}

func TestImage(t *testing.T) {
	body, err := os.ReadFile("1.png")
	if err != nil {
		panic(err)
	}

	hash := md5.New()
	hash.Write(body)
	md5Str := hex.EncodeToString(hash.Sum(nil))
	base64Str := base64.StdEncoding.EncodeToString(body)

	bot := New("98fcbeba-ce84-451b-8b6d-ae162d769fd6")
	msg := &ImageMessage{
		Base64: base64Str,
		Md5:    md5Str,
	}
	if err := bot.Send(msg); err != nil {
		t.Error(err)
	}
}

func TestNews(t *testing.T) {
	bot := New("98fcbeba-ce84-451b-8b6d-ae162d769fd6")
	msg := &NewsMessage{}
	msg.AddArticle(NewsMessageArticle{
		Title:  "test",
		URL:    "https://github.com/duke-git/lancet/blob/main/README_zh-CN.md",
		PicURL: "https://picb9.photophoto.cn/39/872/39872639_1.jpg",
	})
	if err := bot.Send(msg); err != nil {
		t.Error(err)
	}
}

func TestFile(t *testing.T) {
	bot := New("98fcbeba-ce84-451b-8b6d-ae162d769fd6")
	mid, err := bot.UploadMedia("go.mod")
	if err != nil {
		panic(err)
	}
	msg := &FileMessage{
		MediaId: mid,
	}
	if err := bot.Send(msg); err != nil {
		t.Error(err)
	}
}
```
