FROM golang:1.15 AS build

WORKDIR /go/src/github.com/ChorusOne/Icarus
COPY . .
RUN make build

FROM debian:buster

RUN apt update && apt install -y jq bash curl

COPY --from=build /go/src/github.com/ChorusOne/Icarus/build/icarus /usr/local/bin/icarus
COPY --from=build /go/src/github.com/ChorusOne/Icarus/build/icarus-api /usr/local/bin/icarus-api

CMD /usr/bin/local/icarus-api
