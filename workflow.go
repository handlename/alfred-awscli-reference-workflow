package wf

import (
	"encoding/xml"
	"io"
	"strings"
)

type SitemapURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func Convert(in io.Reader, out io.Writer) error {
	dec := xml.NewDecoder(in)

	for {
		t, err := dec.Token()
		if err != nil && err != io.EOF {
			return err
		}

		if t == nil {
			break
		}

		loc := ""

		switch elem := t.(type) {
		case xml.StartElement:
			if elem.Name.Local == "url" {
				// extract url

				var su SitemapURL
				if err := dec.DecodeElement(&su, &elem); err != nil {
					panic(err)
				}

				loc = su.Loc
			}
		default:
			continue
		}

		// convert url to command list

		cmd := locToCommand(loc)
		if cmd == "" {
			continue
		}

		// output

		_, err = out.Write([]byte(cmd + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func locToCommand(loc string) string {
	loc = strings.Replace(loc, ".html", "", -1)
	parts := strings.Split(loc, "/")

	if len(parts) < 8 {
		return ""
	}

	if parts[7] == "index" {
		return ""
	}

	return strings.Join(parts[6:], " ")
}
