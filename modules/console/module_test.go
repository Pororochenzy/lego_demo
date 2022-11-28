package console

import (
	"fmt"
	"testing"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func Test_Bluemonday(t *testing.T) {
	input := []byte("### topgoer.com是个不错的go文档网站")
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Println(string(html))
}
