FROM golang:alpine
ARG PATH="$PATH:$(go env GOPATH)/bin"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/main ./cmd/app 
EXPOSE 8080
CMD ["./build/main"]
