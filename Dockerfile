FROM golang:1.5
COPY ./ /go/src/github.com/hibooboo2/ultimateBravery
WORKDIR /go/src/github.com/hibooboo2/ultimateBravery
EXPOSE 8000
CMD /go/src/github.com/hibooboo2/ultimateBravery/scripts/run
