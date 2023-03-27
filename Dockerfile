FROM --platform=linux/arm64 arm64v8/golang:1.12.0-alpine3.9
RUN mkdir /opt/app
ADD . /opt/app
WORKDIR /opt/app
RUN go build -o main .
CMD ["/opt/app/main"]
