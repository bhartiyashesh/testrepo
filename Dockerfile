FROM golang:1.24-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o server main.go

FROM alpine:3.19
COPY --from=build /app/server /server
ENV PORT=3000
EXPOSE 3000
CMD ["/server"]
