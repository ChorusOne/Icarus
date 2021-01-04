FROM golang:1.15-alpine

COPY . .
RUN make build
