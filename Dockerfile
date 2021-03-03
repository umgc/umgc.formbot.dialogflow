FROM golang

WORKDIR /app

ADD Makefile /app
ADD go.mod /app
ADD go.sum /app
ADD formscriber /app/formscriber

RUN go mod download

EXPOSE 80
EXPOSE 443

CMD make