FROM golang:1.5
COPY ./ /go/src/github.com/hibooboo2/ultimateBravery
WORKDIR /source
EXPOSE 8000
CMD /go/src/github.com/hibooboo2/ultimateBravery
