FROM micro/micro

RUN apk add --no-cache git make musl-dev

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir /opt/go
RUN mkdir /opt/go/modules

RUN apk add protoc

COPY go.mod /opt

WORKDIR /opt
RUN go mod download

WORKDIR /opt/services

ENTRYPOINT ["/micro", "server", "--address=10003"]