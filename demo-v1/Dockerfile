FROM golang:alpine AS build-env
ADD . /src
RUN apk add --no-cache git gcc musl-dev
RUN cd /src && go build -o app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app
