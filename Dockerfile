FROM golang:1.14-alpine

ARG UID=1000

RUN adduser \
  --disabled-password \
  --gecos "" \
  --no-create-home \
  --uid $UID \
  $UID

USER $USER

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
COPY ./ ./

EXPOSE 8080
CMD ["CompileDaemon", "-command=./Legitly-Backend", "-exclude='*.md'", "-log-prefix=false", "-graceful-kill=true"]
