FROM golang:1.14-alpine3.11
WORKDIR /var/app
COPY ./start.server.sh /bin/start.server.sh

RUN apk add --update make
RUN apk add --update git

ENTRYPOINT [ "sh","/bin/start.server.sh" ]