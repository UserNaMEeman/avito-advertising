FROM golang:alpine3.17 as builder
ADD . /src/proj
WORKDIR /src/proj
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o web cmd/main.go

FROM alpine
# WORKDIR /app
COPY --from=builder /src/proj/web /usr/bin
EXPOSE 8080
CMD [ "web" ]