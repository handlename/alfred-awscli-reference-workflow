package wf_test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/handlename/alfred-awscli-reference-workflow"
	alfred "github.com/ruedap/go-alfred"
)

type Response struct {
	alfred.Response
	Items []alfred.ResponseItem `xml:"item"`
}

func TestSearchForOneKeyword(t *testing.T) {
	src := &bytes.Buffer{}
	src.WriteString("acm delete-certificate\n")
	src.WriteString("ec2 wait bundle-task-complete\n")

	out := &bytes.Buffer{}

	err := wf.Search(src, []string{"delete"}, out)
	if err != nil {
		t.Fatalf("failed to search: %s", err.Error())
	}

	res := Response{}
	body, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatalf("failed to read xml body: %s", err)
	}

	if err = xml.Unmarshal(body, &res); err != nil {
		t.Logf("xml: %s", out.String())
		t.Fatalf("failed to decode response: %s", err.Error())
	}

	if num := len(res.Items); num != 1 {
		t.Errorf("invalid items count: %d", num)
	}

	if title := res.Items[0].Title; title != "acm delete-certificate" {
		t.Errorf("unexpected title: %s", title)
	}

	if subtitle := res.Items[0].Subtitle; subtitle != "http://docs.aws.amazon.com/cli/latest/reference/acm/delete-certificate.html" {
		t.Errorf("unexpected subtitle: %s", subtitle)
	}

	if arg := res.Items[0].Arg; arg != "http://docs.aws.amazon.com/cli/latest/reference/acm/delete-certificate.html" {
		t.Errorf("unexpected arg: %s", arg)
	}
}

func TestSearchForMultipleKeywords(t *testing.T) {

}
