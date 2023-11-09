FROM golang:1.20

RUN apt-get update && apt-get install -y

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/api/

EXPOSE 4000
EXPOSE 5432

CMD ["./main"]
