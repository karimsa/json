SRC	=	main.go
TARGET	=	bin/json

all: $(TARGET)

$(TARGET):
	go build -o $(TARGET) $(SRC)

clean:
	rm -f $(TARGET)
