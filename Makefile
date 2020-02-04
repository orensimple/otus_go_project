# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SOURCE_NAME=./api/main.go
BINARY_NAME=bin/api

all: deps gen build test
deps:
	$(GOGET) go.uber.org/zap
	$(GOGET) go.uber.org/zap/zapcore
	$(GOGET) github.com/spf13/cobra
	$(GOGET) github.com/spf13/viper
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) gopkg.in/natefinch/lumberjack.v2
build:
	$(GOBUILD) -o $(BINARY_NAME) $(SOURCE_NAME)
test:
	$(GOTEST) -v internal/domain/services/choice_test.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)  http_server --config ./config
