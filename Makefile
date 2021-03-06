VERSION=0.0.1
FILE_COMMAND=columbus
PATH_BUILD=$(HOME)/go-builds/$(FILE_COMMAND)
FILE_ARCH=darwin_amd64

clean: 
	@rm -rf $(PATH_BUILD)/$(VERSION)

build: clean
	@dep ensure
	@$(GOPATH)/bin/goxc \
	  -bc="darwin,amd64" \
	  -pv=$(VERSION) \
	  -d=$(PATH_BUILD) \
	  -build-ldflags "-X main.VERSION=$(VERSION)"

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)/$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '/usr/local/bin/$(FILE_COMMAND)'
