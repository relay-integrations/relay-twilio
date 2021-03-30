FROM golang:1.13-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -o "/build/relay-send-sms" "./cmd/send-sms"

FROM alpine:3
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder "/build/relay-send-sms" "/usr/bin/relay-send-sms"
ENTRYPOINT []
CMD ["/usr/bin/relay-send-sms"]

LABEL "org.opencontainers.image.title"="Send an SMS using Twilio"
LABEL "org.opencontainers.image.description"="Use Twilio to deliver an SMS message to a recipient"
LABEL "sh.relay.sdk.version"="v1"
