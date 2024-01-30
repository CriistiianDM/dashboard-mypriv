#Golang image
FROM golang:1.19.8

# Set the working directory
WORKDIR /backend

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o mypriv_go ./src/main.go

# Expose the port on which the Go application will run
EXPOSE 4700

# Command to run the Go application
CMD ["./mypriv_go"]
