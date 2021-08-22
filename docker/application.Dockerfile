# Build stage
FROM golang:alpine AS build-env

#RUN apk --no-cache add build-base git mercurial gcc
ADD ./app/ /app
WORKDIR /app
RUN go build -o main.bin cmd/main.go

# Final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/main.bin /app/

ENTRYPOINT ["./main.bin"]