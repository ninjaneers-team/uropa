FROM golang:1.13.4 AS build
WORKDIR /uropa
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD . .
ARG COMMIT
ARG TAG
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o uropa \
      -ldflags "-s -w -X github.com/ninjaneers-team/uropa/cmd.VERSION=$TAG -X github.com/ninjaneers-team/uropa/cmd.COMMIT=$COMMIT"

FROM alpine:3.10
RUN adduser --disabled-password --gecos "" uropauser
RUN apk --no-cache add ca-certificates
USER uropauser
COPY --from=build /uropa/uropa /usr/local/bin
ENTRYPOINT ["uropa"]
