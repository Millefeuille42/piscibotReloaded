FROM golang:1.16.0-alpine3.13

ADD ./checker /checker

WORKDIR /checker

RUN go install .

CMD ["APICheck"]