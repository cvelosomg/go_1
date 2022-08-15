#build stage
FROM golang:latest

RUN mkdir /build
WORKDIR /build

#RUN export GO1111MODULE=on
#RUN go clean -modcache
#RUN go install github.com/cvelosomg/go_1/main@latest
#RUN cd /build && git close https://github.com/cvelosomg/go_1.git

ADD . /build

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]

#ENTRYPOINT [ "/build/GoWebAPI/main/main" ]