FROM golang
ADD main.go .
RUN go get -u github.com/gorilla/mux
RUN go build main.go
ENTRYPOINT [ "./main" ]
