BIN_DIR = ./bin

all:
	go build -o $(BIN_DIR)/wol .

clean:
	rm $(BIN_DIR)/wol

.PHONY: clean