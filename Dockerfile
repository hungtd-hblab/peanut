FROM golang:1.19-buster AS development
WORKDIR /peanut
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

FROM golang:1.19-buster AS build
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
WORKDIR /peanut
ENV CGO_ENABLED=0
ENV GIN_MODE=release
ENV PEANUT_ENV=production
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -tags timetzdata -mod=readonly -v -o peanut
CMD ["./server/peanut"]
