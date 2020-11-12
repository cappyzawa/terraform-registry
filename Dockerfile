FROM golang:1 as builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV CGO_ENABLED 0

RUN go build -o /assets/terraform-registry

FROM alpine:edge
RUN apk add --no-cache bash tzdata ca-certificates unzip zip gzip tar git
COPY --from=builder /assets/terraform-registry .
COPY --from=builder /src/testdata/config.yaml /testdata/config.yaml
RUN chmod +x ./terraform-registry
CMD [ "./terraform-registry"]
