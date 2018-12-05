FROM golang:alpine as build
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux

# gcc and musl-dev required for cgo
RUN apk update && apk upgrade && \
    apk add --no-cache \
    git \
    gcc \ 
    musl-dev

WORKDIR /go/src/app

COPY . .

RUN go build -a -installsuffix cgo -o /app

FROM alpine:latest
COPY --from=build --chown=nobody:nobody /app /app
USER nobody
CMD [ "/app" ]