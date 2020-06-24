GOBUILD=go build
NAME=61231VCH41V

all: build

.PHONY: run
run:
	make build
	./$(NAME).out

.PHONY: build
build:
	CGO_LDFLAGS="`mecab-config --libs`" CGO_CFLAGS="-I`mecab-config --inc-dir`" $(GOBUILD) -o ./$(NAME).out -v

.PHONY: install
install:
	make build
	install -s $(NAME).out /usr/local/bin/$(NAME)
