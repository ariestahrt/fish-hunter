FROM golang:alpine

WORKDIR /app

COPY . /app

RUN apk add --update --no-cache curl
RUN apk add --update --no-cache p7zip

RUN export CRON_HEADER_KEY=$(cat .env | grep CRON_HEADER_VALUE | awk -F'=' '{print $2}')
RUN echo "*/5 * * * *" curl --location --request GET 'https://0.0.0.0:443/api/v1/urls/fetch/openphish' --header "'X-Cron-Key: $CRON_HEADER_KEY'" --insecure >> /var/spool/cron/crontabs/root

RUN go mod download
RUN go build -o dist

EXPOSE 443

ENTRYPOINT [ "./dist" ]
