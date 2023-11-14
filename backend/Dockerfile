FROM golang:1.21

WORKDIR /app

# Copy the go.mod and go.sum files first and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application to a binary named "app" (or any name you prefer)
RUN go build -o backend

# The application will run on port 8080
EXPOSE 8080

# Run the compiled binary
CMD [ "./backend" ]
