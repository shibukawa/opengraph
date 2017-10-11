// Code generated by hero.
// source: /Users/shibu/go/src/github.com/shibukawa/opengraph/opengraphtest/sample.html
// DO NOT EDIT!
package opengraphtest

import (
	"bytes"

	"github.com/shibukawa/opengraph"
)

func WriteHTML(og opengraph.OpenGraph, buffer *bytes.Buffer) {
	buffer.WriteString(`
<!DOCTYPE html>
<html lang="en">
<head `)
	buffer.WriteString(opengraph.NameSpace)
	buffer.WriteString(` >
    <meta charset="UTF-8">
    `)
	buffer.WriteString(og.Write())
	buffer.WriteString(`
    <title>Title</title>
</head>
<body>

</body>
</html>
`)

}
