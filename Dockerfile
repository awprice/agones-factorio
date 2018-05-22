FROM golang:1.10 as builder
WORKDIR /go/src/github.com/awprice/agones-factorio
ADD main.go main.go
RUN go get
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o sdk main.go

FROM dtandersen/factorio:stable
COPY --from=builder /go/src/github.com/awprice/agones-factorio/sdk /bin/sdk
ADD run.sh /run.sh
ENTRYPOINT ["/run.sh"]