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
    
    mkdir -p /var/db
    sqlite3 /var/db/61231VCH41V.sqlite3 | "CREATE TABLE dictionary (id INTEGER PRIMARY KEY, string1 TEXT, string2 TEXT, string3 TEXT);"

    cp ./systemd/*.service /etc/systemd/system/
    cp ./systemd/*.timer   /etc/systemd/system/
    cp ./systemd/*_env     /etc/sysconfig/

    echo 'If necessary, edit /etc/sysconfig/61231VCH41V_env'
