# Alfred AWS CLI Reference Workflow

This is a workflow for [Alfred](http://www.alfredapp.com/).
It searches commands supported by [AWS CLI](https://aws.amazon.com/cli/?nc1=h_ls) and open reference or copy command to clipboard.

## Example

![Searching commands](https://cloud.githubusercontent.com/assets/115636/25463718/14aea92e-2b33-11e7-8c95-15f11a5178a8.png)

## Installation

Download [awscli-reference-x.x.x.alfredworkflow](https://github.com/handlename/alfred-awscli-reference-workflow/releases) to your computer and double click it.

## How to use it

The default invoking keyword is `ac`.
Open Alfred command window, and type `ac {query}`.
You can give multiple keywords separated by white spaces as `{query}`.

For example:

```
ac ec2 describe
```

means, "Search commands which including keyword `ec2` and `describe`."

Choose one in candidates and type `Enter`, command reference will be open in your default web browser.
To type `Cmd`+`Enter` instead of `Enter`, you can copy command to clipboard.

## candidates.txt

This workflow search commands from candidates.txt included in .alfredworkflow file.

```
$ head candidates.txt
acm delete-certificate
acm describe-certificate
acm get-certificate
acm list-certificates
acm request-certificate
acm resend-validation-email
apigateway create-api-key
apigateway create-authorizer
apigateway create-base-path-mapping
apigateway create-deployment
```

You can update candidates by edit the file.

To reveal it:

1. Open Alfred perference
1. Click "Workflows" tab
1. Open context menu on "AWS CLI Reference" in left pane
1. Click "Open in Filder"

Original file has been generated by `convert` command.

```
$ cd $GOPATH/src/github.com/handlename/alfred-awscli-reference-workflow
$ go run cmd/convert/main.go -src-path sitemap.xml > candidates.txt
```

Of course, you must setup environment to build programs written in Go.

You can get `sitemap.xml` by make documentation for AWS CLI.

```
$ git clone git@github.com:aws/aws-cli
$ cd aws-cli/doc
$ make html
$ ls build/html/sitemap.xml
build/html/sitemap.xml
```

## Licence

MIT

## Author

[handlename](https://github.com/handlename)
