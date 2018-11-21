FROM golang:1.11 AS base

RUN go get github.com/golang/dep/cmd/dep

# Get Dependencies
RUN mkdir -p /go/src/github.com/thethan/team_picker
WORKDIR /go/src/github.com/thethan/team_picker/

COPY Gopkg.lock Gopkg.toml /go/src/github.com/thethan/team_picker/
RUN dep ensure -vendor-only

COPY . /go/src/github.com/thethan/team_picker
RUN make test
RUN go build main.go


#RUN GOARCH=wasm GOOS=js go build -o lib.wasm lib.go

FROM golang:1.11
WORKDIR /go/src/github.com/thethan/team_picker/

ENV DATA_LOCATION "data/"
COPY data data
COPY --from=base /go/src/github.com/thethan/team_picker/main main
EXPOSE 80
EXPOSE 8080

CMD ["./main"]