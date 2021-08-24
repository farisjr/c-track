FROM golang:1.16

RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o app

CMD /app/app