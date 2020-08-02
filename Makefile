OUT := bin/httpd
PKG := github.com/developersam1995/cabs-service/cmd/http

build:
	go build -i -v -o ${OUT} ${PKG}
 
test:
	go test -cover ./...