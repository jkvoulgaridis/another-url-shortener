FROM golang:1.21.5
WORKDIR /

COPY ./backend .

RUN go mod download

RUN go build
RUN chmod +x url-short