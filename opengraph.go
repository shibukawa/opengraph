package opengraph

//go:generate hero -pkgname opengraph

import (
	"path"
	"strings"
	"time"
)

const NameSpace = `prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# article: http://ogp.me/ns/article#"`

func NewProfile(domain, author, orgName, orgLogoURL, siteName, twitterID, facebookAppID string) *Profile {
	if !strings.HasPrefix(orgLogoURL, "http") {
		orgLogoURL = path.Join(domain, orgLogoURL)
	}
	return &Profile{
		domain:        domain,
		author:        author,
		orgName:       orgName,
		orgLogoURL:    orgLogoURL,
		siteName:      siteName,
		twitterID:     twitterID,
		facebookAppID: facebookAppID,
	}
}

type Profile struct {
	domain        string
	author        string
	orgName       string
	orgLogoURL    string
	siteName      string
	twitterID     string
	facebookAppID string
}

type OpenGraph interface {
	Write() string
}

func (p *Profile) Article(apath, title, description, imageUrl string, pubTime time.Time, modTime ...time.Time) *Article {
	if !strings.HasPrefix(imageUrl, "http") {
		imageUrl = path.Join(p.domain, imageUrl)
	}
	result := &Article{
		profile:     p,
		url:         path.Join(p.domain, apath),
		imageUrl:    imageUrl,
		title:       title,
		description: description,
		pubTime:     pubTime,
	}
	if len(modTime) > 0 {
		result.modTime = modTime[0]
	} else {
		result.modTime = pubTime
	}
	return result
}

type Article struct {
	profile     *Profile
	url         string
	imageUrl    string
	title       string
	description string
	pubTime     time.Time
	modTime     time.Time
	Creator     string
}
