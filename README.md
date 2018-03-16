# widgets-api

This project was developed using golang 1.10 version

## How to build and run

Be sure you have golang and dep(https://golang.github.io/dep/) installed into your system
```
Clone this project into your Go src directory
git clone https://github.com/mariouzae/widgets-api $GOPATH/src

cd $GOPATH/src/widgets-api && dep ensure
cd $GOPATH/src/widgets-api && go install
cd $GOPATH/bin && ./widgets-api
```

Type in your browser

https://localhost:4000

**WARNING**: This project use an invalid SSL certificate which must be accepted. Most of browsers like Chrome leave you to proceed anyway with non valids certificates.
