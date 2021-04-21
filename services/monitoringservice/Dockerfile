FROM golang:latest

ENV GOPRIVATE github.com/VladyslavLukyanenko/GopherAlert

WORKDIR /go/src/app
COPY . .

ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/VladyslavLukyanenko/".insteadOf "https://github.com/VladyslavLukyanenko/"

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["MonitoringService", "run"]