FROM golang:1.20 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o turnttable

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /app/turnttable /turnttable
ENTRYPOINT [ "/turnttable" ]