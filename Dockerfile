FROM golang
ADD webpage.go .
RUN go get -u github.com/gorilla/mux
RUN go build webpage.go
ENTRYPOINT [ "./webpage" ]
