#build stage
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO1111MODULE=on
RUN go get github.com/adminTurnedDevOps/GoWebAPI/main
RUN cd /build && git close https://github.com/AdminTurnedDevOps/GoWebAPI.git

EXPOSE 8080

ENTRYPOINT [ "/build/GoWebAPI/main/main" ]