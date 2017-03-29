package wf

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	alfred "github.com/ruedap/go-alfred"
)

type SitemapURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func Search(src io.Reader, keywords []string, out io.Writer) error {
	cmds, err := filter(src, keywords)
	if err != nil {
		return err
	}

	err = writeAlfredXML(cmds, out)
	if err != nil {
		return err
	}

	return nil
}

func filter(src io.Reader, keywords []string) (cmds []string, err error) {
	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		cmd := scanner.Text()
		match := true

		for _, keyword := range keywords {
			if !strings.Contains(cmd, keyword) {
				match = false
				break
			}
		}

		if match {
			cmds = append(cmds, cmd)
		}
	}

	return cmds, nil
}

func writeAlfredXML(cmds []string, out io.Writer) error {
	res := alfred.NewResponse()

	for _, cmd := range cmds {
		u := cmdToURL(cmd)

		item := alfred.ResponseItem{
			Valid:    true,
			UID:      u,
			Arg:      u,
			Title:    cmd,
			Subtitle: u,
		}

		res.AddItem(&item)
	}

	body, err := res.ToXML()
	if err != nil {
		return err
	}

	fmt.Fprint(out, body)

	return nil
}

func cmdToURL(cmd string) string {
	parts := strings.Split(cmd, " ")

	return fmt.Sprintf("http://docs.aws.amazon.com/cli/latest/reference/%s.html", strings.Join(parts, "/"))
}

// Convert converts sitemap.xml to command list.
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
