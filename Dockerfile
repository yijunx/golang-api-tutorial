FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# COPY *.go ./
ADD . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build ./cmd/api/main.go


# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM ubuntu:latest

WORKDIR /

COPY --from=build-stage /app/main /

EXPOSE 8000

ENTRYPOINT ["./main"]


