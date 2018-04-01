FROM golang:alpine
LABEL maintainer="Mario Uzae" email="mariouzae@gmail.com"
# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml /go/src/widgets-api/
WORKDIR /go/src/widgets-api/
# Install library dependencies
RUN dep ensure -vendor-only
# Copy all project and build it
# This layer will be rebuilt when ever a file has changed in the project directory
COPY . /go/src/widgets-api/
RUN go install
WORKDIR /go/bin/
ENTRYPOINT ["/go/bin/widgets-api"]
EXPOSE 4000 