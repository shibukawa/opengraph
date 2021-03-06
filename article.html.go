// Code generated by hero.
// source: /Users/shibu/go/src/github.com/shibukawa/opengraph/article.html
// DO NOT EDIT!
package opengraph

import (
	"bytes"
	"time"

	"github.com/shiyanhui/hero"
)

func (a *Article) Write() string {
	var buffer bytes.Buffer
	a.write(&buffer)
	return buffer.String()
}
func (a *Article) write(buffer *bytes.Buffer) {
	buffer.WriteString(`
<meta property="og:type" content="artcle" />
<meta content="`)
	hero.EscapeHTML(a.description, buffer)
	buffer.WriteString(`" name="description" />
<meta content="summary" name="twitter:card" />
<meta content="`)
	hero.EscapeHTML(a.profile.twitterID, buffer)
	buffer.WriteString(`" name="twitter:site" />
`)
	if a.Creator != "" {
		buffer.WriteString(`<meta content="`)
		hero.EscapeHTML(a.Creator, buffer)
		buffer.WriteString(`" name="twitter:creator" />`)
	}
	buffer.WriteString(`
<meta content="`)
	hero.EscapeHTML(a.title, buffer)
	buffer.WriteString(`" property="og:title" />
<meta content="`)
	hero.EscapeHTML(a.url, buffer)
	buffer.WriteString(`" property="og:url" />
<meta content="`)
	hero.EscapeHTML(a.imageUrl, buffer)
	buffer.WriteString(`" property="og:image" />
<meta content="`)
	hero.EscapeHTML(a.description, buffer)
	buffer.WriteString(`" property="og:description" />
<meta content="`)
	hero.EscapeHTML(a.profile.siteName, buffer)
	buffer.WriteString(`" property="og:site_name" />
<meta content="`)
	hero.EscapeHTML(a.profile.facebookAppID, buffer)
	buffer.WriteString(`" property="fb:admins" />
<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "NewsArticle",
  "mainEntityOfPage": {
    "@type": "WebPage",
    "@id": "`)
	hero.EscapeHTML(a.url, buffer)
	buffer.WriteString(`"
  },
  "headline": "`)
	hero.EscapeHTML(a.title, buffer)
	buffer.WriteString(`",
  "image": [
    "`)
	hero.EscapeHTML(a.imageUrl, buffer)
	buffer.WriteString(`"
   ],
  "datePublished": "`)
	hero.EscapeHTML(a.pubTime.Format(time.RFC3339), buffer)
	buffer.WriteString(`",
  "dateModified": "`)
	hero.EscapeHTML(a.modTime.Format(time.RFC3339), buffer)
	buffer.WriteString(`",
  "author": {
    `)
	if a.profile.author == "" {
		buffer.WriteString(`
    "@type": "Organization",
    "name": "`)
		hero.EscapeHTML(a.profile.orgName, buffer)
		buffer.WriteString(`",
    "logo": {
      "@type": "ImageObject",
      "url": "`)
		hero.EscapeHTML(a.profile.orgLogoURL, buffer)
		buffer.WriteString(`"
    }
    `)
	} else {
		buffer.WriteString(`
      "@type": "Person",
      "name": "`)
		hero.EscapeHTML(a.profile.author, buffer)
		buffer.WriteString(`"
    `)
	}
	buffer.WriteString(`
  },
   "publisher": {
    "@type": "Organization",
    "name": "`)
	hero.EscapeHTML(a.profile.orgName, buffer)
	buffer.WriteString(`",
    "logo": {
      "@type": "ImageObject",
      "url": "`)
	hero.EscapeHTML(a.profile.orgLogoURL, buffer)
	buffer.WriteString(`"
    }
  },
  "description": "`)
	hero.EscapeHTML(a.description, buffer)
	buffer.WriteString(`"
}
</script>`)

}
