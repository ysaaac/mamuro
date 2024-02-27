FROM golang:1.22.0-alpine

# Creates and Sets the working directory in the container
WORKDIR /app

# Copy all files to app folder
COPY ./api /app

# Install dependencies
RUN go mod download

# Starts building go app
RUN go build -o main .

# Run
CMD ["./main"]