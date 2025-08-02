FROM golang:1.24.3 AS builder

ARG VERSION
ARG VCS_URL
ARG VCS_REF
ARG VCS_BRANCH
ARG BUILD_DATE

ENV GOPATH=/go
ENV GOOS=linux
ENV CGO_ENABLED=0

ARG GO_PROXY
RUN go env -w GOPROXY=${GO_PROXY}

WORKDIR /go/src/github.com/hatena/ipdrawer
COPY . .
RUN make

RUN cp /go/src/github.com/hatena/ipdrawer/ipdrawer /bin/ipdrawer
RUN /bin/ls -l /bin/ipdrawer

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/hatena/ipdrawer/ipdrawer /bin/ipdrawer

ENTRYPOINT ["ipdrawer"]
