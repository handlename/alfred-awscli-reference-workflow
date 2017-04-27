VERSION=${shell cat VERSION}
CMD=cmd/workflow/workflow
WORKFLOW=awscli-reference-$(VERSION).alfredworkflow

dist/$(WORKFLOW): $(CMD) candidates.txt icon.png info.plist
	zip -jv $@ $^

$(CMD): *.go
	cd cmd/workflow; go build -v -ldflags "-X main.version=$(VERSION)"

.PHONY: clean
clean:
	-rm $(CMD) dist/$(WORKFLOW)

.PHONY: ship
ship:
	ghr $(VERSION) dist/$(WORKFLOW)
