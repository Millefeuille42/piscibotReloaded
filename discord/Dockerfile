FROM golang:1.16.0-alpine3.13

ADD ./segbot /segbot/
WORKDIR /segbot

RUN go install .

CMD segbot ./dev.env