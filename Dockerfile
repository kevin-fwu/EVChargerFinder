FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_HOME /go/src/EVChargerFinder
WORKDIR "$APP_HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o evchargerfinder

FROM registry.semaphoreci.com/golang:1.18

ENV APP_HOME /go/src/EVChargerFinder
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/evchargerfinder $APP_HOME

EXPOSE 8080
CMD ["./evchargerfinder" "-conf=evchargerfinder.json"]
