FROM golang:1.22 AS build
WORKDIR /src

COPY go.mod main.go ./

RUN go build -o /out/ainews ./main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /out/ainews /app/ainews
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/ainews"]
