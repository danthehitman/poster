#FROM golang:onbuild
#EXPOSE 8080

FROM golang:1.6
ENV GOPATH /go:/go/src/poster/src
ADD . /go/src/poster
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/kardianos/osext
RUN go install poster
ENTRYPOINT /go/bin/poster
EXPOSE 8080
