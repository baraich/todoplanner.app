# Using the official golang image from Docker hub.
FROM golang:1.23

# Setting the working directory for this image to /app.
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copying the go files.
COPY ./ ./

# Building the project
RUN CGO_ENABLED=0 GOOS=linux go build -o ssh ./cmd/main.go

# Exposing 2222/tcp
EXPOSE 2222

# Running the project.
ENTRYPOINT [ "./ssh" ]
