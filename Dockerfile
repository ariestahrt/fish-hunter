FROM golang:alpine

WORKDIR /app

RUN apk add --update --no-cache curl
RUN apk add --update --no-cache p7zip
RUN apk add --update --no-cache nano

EXPOSE 443

ENTRYPOINT [ "/bin/sh" ]
