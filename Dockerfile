FROM golang:1.25.4-alpine

WORKDIR /app
ARG bin_to_build
ARG port_to_expose
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o svr cmd/${bin_to_build}/main.go

EXPOSE ${port_to_expose}
CMD [ "./svr" ]