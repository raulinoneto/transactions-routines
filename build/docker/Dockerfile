FROM golang:1.14-alpine3.11
WORKDIR /var/app
COPY ./start.server.sh /bin/start.server.sh

RUN apk add --update make
RUN apk add --update git
RUN apk add --update npm
RUN apk add --update musl-dev

RUN npm -g install newman

ENTRYPOINT [ "sh","/bin/start.server.sh" ]