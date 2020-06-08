package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/mdigger/goldmark-stats"
	"github.com/mdigger/goldmark-toc"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"

	"github.com/polaris1119/go-demo/goldmark/mention"
)

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
}

func demo1() {
	source, err := ioutil.ReadFile("guide.md")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("guide1.html")
	if err != nil {
		panic(err)
	}

	err = goldmark.Convert(source, f)
	if err != nil {
		panic(err)
	}
}

func demo2() {
	source, err := ioutil.ReadFile("guide.md")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("guide2.html")
	if err != nil {
		panic(err)
	}

	// 自定义解析器
	markdown := goldmark.New(
		// 支持 GFM
		goldmark.WithExtensions(extension.GFM),
	)

	err = markdown.Convert(source, f)
	if err != nil {
		panic(err)
	}
}

func demo3() {
	source, err := ioutil.ReadFile("guide.md")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("guide3.html")
	if err != nil {
		panic(err)
	}

	// 自定义解析器
	markdown := goldmark.New(
		// 支持 GFM
		goldmark.WithExtensions(extension.GFM),
		// 语法高亮
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
	)

	err = markdown.Convert(source, f)
	if err != nil {
		panic(err)
	}
}

func demo4() {
	source, err := ioutil.ReadFile("guide.md")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("guide4.html")
	if err != nil {
		panic(err)
	}

	// 自定义解析器
	markdown := goldmark.New(
		// 支持 GFM
		goldmark.WithExtensions(extension.GFM),
		// 语法高亮
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
		// 支持 @
		goldmark.WithExtensions(mention.Mention),
	)

	convertFunc := toc.Markdown(markdown)
	headers, err := convertFunc(source, f)

	for _, header := range headers {
		fmt.Printf("%+v\n", header)
	}

	// stats
	doc := goldmark.DefaultParser().Parse(text.NewReader(source))
	info := stats.New(doc, source)

	fmt.Printf("words: %d, unique: %d, chars: %d, reading time: %v\n",
		info.Words, info.Unique(), info.Chars, info.Duration(400))
}
