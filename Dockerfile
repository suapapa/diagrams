FROM golang:1.19 as builder
WORKDIR /src
COPY go.mod .
COPY . .
WORKDIR /src/container
RUN CGO_ENABLED=0 go build -o sandbox

FROM python:alpine
RUN apk add --update --no-cache graphviz \
    font-bitstream-type1 ghostscript-fonts ttf-cantarell
RUN pip install diagrams
WORKDIR /diagrams
COPY container/listup_nodes.py .
COPY --from=builder /src/container/sandbox .
ENTRYPOINT ["/diagrams/sandbox"]
CMD [""]
