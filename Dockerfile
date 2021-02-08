FROM golang:1.15 AS build

# Set necessary environmet variables needed for our image
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . /app
WORKDIR /app
RUN go get -d
RUN go build -o main .


FROM alpine:latest AS runtime

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY --from=build /app/main ./
# Export necessary port
EXPOSE 80

# Command to run when starting the container
CMD ["./main"]