#Golang image
FROM golang:1.19.8

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o mypriv_go ./src/main.go

EXPOSE 4700

CMD ["./mypriv_go"]
