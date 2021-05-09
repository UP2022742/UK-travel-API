FROM golang:1.16.4

WORKDIR /

COPY . .
RUN go get -d github.com/gorilla/mux && \
    go get -d github.com/inconshreveable/log15

CMD ["./main.exe"]