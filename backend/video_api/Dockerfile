FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

EXPOSE 3000

CMD ["go", "run", "main.go"]

