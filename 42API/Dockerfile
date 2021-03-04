FROM golang:1.16.0-alpine3.13

ADD ./api /api

WORKDIR /api

RUN go install .

CMD ["42API"]
