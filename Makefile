BINDIR := $(PWD)/bin
PATH:=$(PWD)/bin:$(PATH)

.PHONY: all
all: textReader sumInts searchString

.PHONY: textReader
textReader: textReader/main.go
	go build -o $(BINDIR)/textReader ./textReader

.PHONY: searchString
searchString: searchString/main.go
	go build -o $(BINDIR)/searchString ./searchString

.PHONY: sumInts
sumInts: sumInts/main.go generate_sums
	go build -o $(BINDIR)/sumInts ./sumInts

.PHONY: generate_sums
generate_sums: bootstrap
	go generate ./sumInts/...

.PHONY: bootstrap
bootstrap: sumInts/sums/main.go
	GOBIN=$(BINDIR) go install ./sumInts/sums

.PHONY: clean
clean:
	rm -rf bin