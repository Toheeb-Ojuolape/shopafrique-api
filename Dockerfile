# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .
# Add your .env files
COPY .env . 

# Build the Golang binary
RUN go build -o main .

# Expose the port your application will listen on
EXPOSE 5000

# Command to run your application
CMD ["./main"]
