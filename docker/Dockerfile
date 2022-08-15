#build stage
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO1111MODULE=on
RUN go install github.com/cvelosomg/go_1/main@latest
RUN cd /build && git close https://github.com/cvelosomg/go_1.git

EXPOSE 8080

ENTRYPOINT [ "/build/GoWebAPI/main/main" ]