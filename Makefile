BIN_DIR = ./bin

all: wol

wol:
	go build -o $(BIN_DIR)/wol .