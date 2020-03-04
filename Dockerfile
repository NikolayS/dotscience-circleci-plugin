FROM golang:1.13.8-alpine3.11 AS build-env

WORKDIR /usr/local/go/src/github.com/dotmesh-io/dotscience-circleci-plugin
COPY . /usr/local/go/src/github.com/dotmesh-io/dotscience-circleci-plugin

RUN apk update && apk upgrade
ENV GO111MODULE=off
RUN cd cmd/ds-circleci-plugin && go install

FROM alpine:latest
LABEL "com.dotscience.dotscience-circleci-plugin"="true"
COPY --from=build-env /usr/local/go/bin/ds-circleci-plugin /bin/ds-circleci-plugin

CMD ["ds-circleci-plugin"]
