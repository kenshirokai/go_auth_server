FROM golang

WORKDIR /app/go

COPY ./app/ ./

ENV http_proxy "http://50.27.190.111:8080"
ENV https_proxy "http://50.27.190.111:8080"
ENV HTTP_PROXY "http://50.27.190.111:8080"
ENV HTTPS_PROXY "http://50.27.190.111:8080"
ENV no_proxy "127.0.0.1,localhost,0.0.0.0"
ENV NO_PROXY "127.0.0.1,localhost,0.0.0.0"

RUN git config --global http.proxy http://50.27.190.111:8080
RUN git config --global https.proxy http://50.27.190.111:8080
RUN go version
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64  \ 
    go build -a -tags netgo -installsuffix netgo ./cmd/main.go 

ENTRYPOINT [ "./main", "-mode=prod" ]