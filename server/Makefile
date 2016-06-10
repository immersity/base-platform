.all:
GOOS = linux
GOFLAGS = -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo
CMD_PKG = main.go
TARGET = bin/base

.PHONY: all build clean

all: $(TARGET)

clean:
	$(RM) $(TARGET)
	$(RM) -r bin

generate:
	go generate ./...

$(TARGET): clean generate
	mkdir -p bin
	GOOS=$(GOOS) go build $(GOFLAGS) -o $(TARGET) $(CMD_PKG)
