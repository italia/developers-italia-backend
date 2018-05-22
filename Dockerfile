# build stage
FROM golang:1.10.0-alpine AS build-env
ARG NAME
ARG PROJECT
ARG VERSION

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git && \
    apk add --no-cache gcc && \
    apk add --no-cache musl-dev

ADD . /go/src/$PROJECT

# Dep ensure. Uncomment if you don't have a ./vendor folder for go deps.
# RUN cd /go/src/$PROJECT && go get -u github.com/golang/dep/cmd/dep && dep ensure

# Compile project
RUN cd /go/src/$PROJECT && go build -ldflags "-X github.com/italia/developers-italia-backend/version.VERSION=${VERSION}" -o $NAME

# final stage
FROM alpine:3.7
ARG NAME
ARG PROJECT

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=build-env /go/src/$PROJECT/$NAME /app/
COPY --from=build-env /go/src/$PROJECT/domains.yml /app/
COPY --from=build-env /go/src/$PROJECT/config.toml /app/
EXPOSE 8081

# ARG values are not allowed in ENTRYPOINT, pass NAME as ENV variable.
ENV NAME=$NAME
RUN chmod +x ./$NAME

ENTRYPOINT ./$NAME crawl
