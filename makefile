GOBUILD=go build
BIN_NAME=61231VCH41V.out

.PHONY: run
run:
	CGO_LDFLAGS="`mecab-config --libs`" CGO_CFLAGS="-I`mecab-config --inc-dir`" $(GOBUILD) -o ./$(BIN_NAME) -v
	./$(BIN_NAME)

.PHONY: build
build:
	CGO_LDFLAGS="`mecab-config --libs`" CGO_CFLAGS="-I`mecab-config --inc-dir`" $(GOBUILD) -o ./$(BIN_NAME) -v
