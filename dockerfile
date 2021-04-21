FROM golang:alpine AS build

RUN apk add --no-cache make git
WORKDIR /go/src/app
COPY . .
RUN make build

FROM alpine

COPY --from=build /go/src/app/bin /
ENTRYPOINT ["/mailcheck"]
