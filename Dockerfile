FROM golang:1.16.3-alpine3.13 AS builder

COPY . /ytmp3/
WORKDIR /ytmp3/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

RUN apk update && apk add curl && apk add ffmpeg

ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

WORKDIR /root/

RUN mkdir music

COPY --from=0 /ytmp3/bin/bot .

CMD ["./bot"]