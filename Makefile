NAME=crack
BUILDPATH=cmd/crack/crack.go
BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknown version")
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'

docker:
	$(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(BUILDPATH)

sha256sum:
	cd $(BINDIR); for file in *; do sha256sum $$file > $$file.sha256; done

clean:
	rm $(BINDIR)/*