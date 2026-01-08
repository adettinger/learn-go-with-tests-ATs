# Make sure to specify the same Go version as the one in the go.mod file.
FROM golang:1.25.4-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o svr cmd/grpcserver/*.go

EXPOSE 50051
CMD [ "./svr" ]