FROM --platform=amd64 golang:latest
RUN mkdir -p /go/src/kuik8srampup/promMetrics
WORKDIR /go/src/kuik8srampup/promMetrics
ADD . ./
# ENV GOPATH=/gopath
# ENV GOROOT=/go
ENV PATH=$PATH:$GOPATH/bin
RUN go install kuik8srampup/promMetrics
EXPOSE 8088
CMD ["promMetrics"]