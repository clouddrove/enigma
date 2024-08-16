ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} as builder
ARG PROGRAM=nothing
ARG VERSION=development

RUN mkdir /src /output

WORKDIR /src

COPY . .
RUN GOBIN=/output make install VERSION=$VERSION
RUN PROGRAM=$(ls /output); echo "#!/bin/sh\nexec '/usr/bin/$PROGRAM' \"\$@\"" > /docker-entrypoint.sh && chmod +x /docker-entrypoint.sh


FROM gcr.io/distroless/base:latest
ARG PROGRAM=nothing

COPY --from=builder /output/${PROGRAM} /
USER 1000
ENTRYPOINT [""]