TARGET := ./app

all:    $(TARGET)

$(TARGET): $(wildcard **/*.go)
	CGO_ENABLED=0 GOOS=linux go build \
		-a \
		-installsuffix cgo \
		-o $(TARGET)
