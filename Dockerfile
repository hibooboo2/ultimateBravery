FROM golang:1.5
COPY ./ /ultimateBravery
WORKDIR /source
EXPOSE 8000
CMD /ultimateBravery/scripts/run
