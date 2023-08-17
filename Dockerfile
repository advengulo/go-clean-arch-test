##
## Build
##
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /usr/src/linux-headers-test/go-srv-test/

# Copy application data into image
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/

##
## Deploy
##
FROM alpine:latest

# Create a directory for data (you can modify this as needed)
RUN mkdir /data

# Copy the built server executable from the builder stage
COPY --from=builder /server ./

# Copy the .env file into the container
COPY .env ./

# Set the startup command for the container
CMD ["./server"]