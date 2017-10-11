package opengraph

import (
	"strings"
	"testing"
	"time"
)

func TestGenerateArticle(t *testing.T) {
	p := NewProfile(
		"https://www.example.com",
		"Yoshiki",
		"Yoshiki", "/image/logo.png",
		"shibu.jp diary", "@shibu_jp", "afdasfasfas")
	a := p.Article("/article", "My Favorite Music", "test", "/image/blog.png", time.Now())

	result := a.Write()
	if !strings.Contains(result, `"headline": "My Favorite Music"`) {
		t.Log(result)
		t.Error("error")
	}
}
