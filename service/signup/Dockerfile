FROM golang:alpine AS builder
WORKDIR /signup
COPY . .
RUN go build

FROM alpine
WORKDIR /signup
COPY --from=builder /signup/signup ./

EXPOSE 3001
ENTRYPOINT ["./signup"]
