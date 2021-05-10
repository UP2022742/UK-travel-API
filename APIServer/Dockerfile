FROM golang:1.15-alpine AS GO_BUILD
WORKDIR /server
COPY . /server

RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/gocolly/colly
RUN go get -d github.com/inconshreveable/log15
RUN go build -o /go/bin/server

FROM alpine:3.10 
COPY --from=GO_BUILD /go/bin/server/ ./
COPY --from=GO_BUILD /server/certs ./certs
EXPOSE 8080 
CMD ./server