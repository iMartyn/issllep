FROM golang:onbuild
RUN git clone https://github.com/iMartyn/issllep.git && \
cd issllep; GOPATH=`pwd` go get github.com/mattn/go-sqlite3; go build proxy.go && \
mv proxy /root/
