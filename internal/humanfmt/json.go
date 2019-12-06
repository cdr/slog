package humanfmt

import (
	"bytes"
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	jlexers "github.com/alecthomas/chroma/lexers/j"
)

// Adapted from https://github.com/alecthomas/chroma/blob/2f5349aa18927368dbec6f8c11608bf61c38b2dd/styles/bw.go#L7
// https://github.com/alecthomas/chroma/blob/2f5349aa18927368dbec6f8c11608bf61c38b2dd/formatters/tty_indexed.go
// https://github.com/alecthomas/chroma/blob/2f5349aa18927368dbec6f8c11608bf61c38b2dd/lexers/j/json.go
var style = chroma.MustNewStyle("slog", chroma.StyleEntries{
	// Magenta.
	chroma.Keyword: "#7f007f",
	// Magenta.
	chroma.Number: "#7f007f",
	// Magenta.
	chroma.Name: "#00007f",
	// Green.
	chroma.String: "#007f00",
})

var jsonLexer = chroma.Coalesce(jlexers.JSON)

func formatJSON(w io.Writer, buf []byte) []byte {
	if !shouldColor(w) {
		return buf
	}

	highlighted, err := colorizeJSON(buf)
	if err != nil {
		println("humanfmt: failed to colorize fields JSON: " + err.Error())
		return buf
	}
	return highlighted
}

func colorizeJSON(buf []byte) ([]byte, error) {
	it, err := jsonLexer.Tokenise(nil, string(buf))
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	err = formatters.TTY8.Format(b, style, it)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
