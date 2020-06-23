.PHONY: run
run:
	export CGO_LDFLAGS="`mecab-config --libs`"
	export CGO_CFLAGS="-I`mecab-config --inc-dir`"
	go run ./main.go

build:
	export CGO_LDFLAGS="`mecab-config --libs`"
	export CGO_CFLAGS="-I`mecab-config --inc-dir`"
	go build -o ./61231VCH41V.out