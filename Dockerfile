FROM ysicing/goa AS gobuild

LABEL maintainer="ysicing <i@ysicing.me>"

WORKDIR /go/src

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod download

COPY . .

WORKDIR /go/src/cmd

RUN go build -o ./rv

FROM ysicing/debian

COPY --from=gobuild /go/src/cmd/rv /bin/

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh /bin/rv

ENTRYPOINT ["/entrypoint.sh", "/bin/rv"]