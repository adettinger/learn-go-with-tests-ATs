# Make sure to specify the same Go version as the one in the go.mod file.
# For example, golang:1.22.1-alpine.
FROM golang:1.25.4-alpine

# Install debugging tools
# RUN apk add --no-cache curl bash

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o svr cmd/httpserver/*.go
EXPOSE 8080

# Add a script that will run your app and keep the container alive
# COPY entrypoint.sh /app/
# RUN chmod +x /app/entrypoint.sh
# CMD ["/app/entrypoint.sh"]

CMD [ "./svr" ]