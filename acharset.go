package acharset

import (
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

var charsetToCharmap = map[string]encoding.Encoding{
	"windows-874":  charmap.Windows874,
	"windows-1250": charmap.Windows1250,
	"windows-1251": charmap.Windows1251,
	"windows-1252": charmap.Windows1252,
	"windows-1253": charmap.Windows1253,
	"windows-1254": charmap.Windows1254,
	"windows-1255": charmap.Windows1255,
	"windows-1256": charmap.Windows1256,
	"windows-1257": charmap.Windows1257,
	"windows-1258": charmap.Windows1258,
	"koi8r":        charmap.KOI8R,
	"koi8u":        charmap.KOI8U,
	"iso-8859-1":   charmap.ISO8859_1,
	"iso-8859-2":   charmap.ISO8859_2,
	"iso-8859-3":   charmap.ISO8859_3,
	"iso-8859-4":   charmap.ISO8859_4,
	"iso-8859-5":   charmap.ISO8859_5,
	"iso-8859-6":   charmap.ISO8859_6,
	"iso-8859-7":   charmap.ISO8859_7,
	"iso-8859-8":   charmap.ISO8859_8,
	"iso-8859-10":  charmap.ISO8859_10,
	"iso-8859-13":  charmap.ISO8859_13,
	"iso-8859-14":  charmap.ISO8859_14,
	"iso-8859-15":  charmap.ISO8859_15,
}

func makeReader(charset string, input io.Reader) (io.Reader, error) {
	if charEncoding, ok := charsetToCharmap[charset]; ok {
		return charEncoding.NewDecoder().Reader(input), nil
	}
	return input, nil
}

func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	return makeReader(charset, input)
}
