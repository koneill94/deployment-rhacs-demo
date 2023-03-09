FROM golang:1.18

WORKDIR /usr/src/app

COPY ./src/hello-world /usr/src/app/hello-world

CMD ["/usr/src/app/hello-world"]
