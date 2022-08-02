FROM golang:1.16-alpine

RUN apk add --no-cache build-base mupdf mupdf-dev
WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download && go build -o ticketsystem main.go

CMD [ "./ticketsystem" ]