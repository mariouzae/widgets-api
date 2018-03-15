FROM golang as compile
RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/mariouzae/widgets-api

FROM scratch
COPY --from=compile /go/bin/widgets-api .
EXPOSE 8080
CMD ["./widgets-api"]