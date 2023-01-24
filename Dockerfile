ARG env

FROM golang:1.19-alpine

ARG env
ENV env=$env

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ .
COPY ./config-${env}.yaml config.yaml

RUN go mod tidy
RUN go build -o privilege 
ENV TZ=Asia/Bangkok

EXPOSE 8000
ENTRYPOINT [ "./privilege" ]
