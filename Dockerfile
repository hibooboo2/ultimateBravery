FROM golang:1.5
COPY ./ /source
WORKDIR /source
EXPOSE 8000
CMD /source/scripts/run
