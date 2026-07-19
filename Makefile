BINARY := qsd77
PREFIX := /usr
BINDIR := $(PREFIX)/bin

.PHONY: build install uninstall clean

build:
	go build -o $(BINARY)

install: build
	install -Dm755 $(BINARY) $(DESTDIR)$(BINDIR)/$(BINARY)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/$(BINARY)

clean:
	rm -f $(BINARY)
