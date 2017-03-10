package wf_test

import (
	"bytes"
	"testing"

	"github.com/handlename/alfred-awscli-reference-workflow"
)

func TestConvertSitemapSuccess(t *testing.T) {
	in := bytes.NewBufferString(`
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>http://docs.aws.amazon.com/cli/latest/index.html</loc>
    </url>
    <url>
        <loc>http://docs.aws.amazon.com/cli/latest/reference/acm/delete-certificate.html</loc>
    </url>
    <url>
        <loc>http://docs.aws.amazon.com/cli/latest/reference/ec2/wait/bundle-task-complete.html</loc>
    </url>
</urlset>
`)
	out := &bytes.Buffer{}

	err := wf.Convert(in, out)
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	expected := []byte(`acm delete-certificate
ec2 wait bundle-task-complete
`)

	if bytes.Compare(expected, out.Bytes()) != 0 {
		t.Fatalf("unexpected out: %s", string(out.Bytes()))
	}
}
