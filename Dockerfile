# Build the goto binary
FROM golang:1.20 AS build

WORKDIR /app

ADD . .

RUN go mod tidy
RUN CGENABLED=0 GOOS=linux go build -o bin/gotodo main.go

EXPOSE 9000

ENTRYPOINT [ "./bin/gotodo" ]

# TODO: Make the multistage build work
# Getting the "exec ./gotodo: no such file or directory"

# Get the binary from the build stage and
# Execute in alpine container
# FROM alpine:3.17

# RUN apk add --no-cache ca-certificates &&\
#     rm -rf /var/lib/cache

# WORKDIR /app

# COPY --from=build /app/bin/gotodo /app/

# EXPOSE 9000

# ENTRYPOINT [ "./gotodo" ]