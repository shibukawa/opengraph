package opengraphtest

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/shibukawa/opengraph"
)

func TestEmbedded(t *testing.T) {
	p := opengraph.NewProfile(
		"https://www.example.com",
		"Yoshiki",
		"Yoshiki", "/image/logo.png",
		"shibu.jp diary", "@shibu_jp", "afdasfasfas")
	article := p.Article(
		"/article", "My Favorite Music",
		"test", "/image/blog.png", time.Now())

	var buffer bytes.Buffer
	WriteHTML(article, &buffer)
	fmt.Println(buffer.String())
}
