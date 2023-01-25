FROM shabykov12/catboost-c-api:1.1.1-golang-1.16.3-streach AS builder

ARG OUTPUT_BINARY=/bin

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOSUMDB=off
ENV GOPROXY=direct

COPY . /app/
WORKDIR /app

RUN go build -o $OUTPUT_BINARY

FROM golang:1.16.3-stretch

ARG CATBOOST_BIN=/app/catboost_bin
ARG OUTPUT_BINARY=/bin

RUN mkdir /app

COPY --from=builder $CATBOOST_BIN $CATBOOST_BIN
COPY --from=builder $OUTPUT_BINARY/ /app/

ENV LD_LIBRARY_PATH=$CATBOOST_BIN

RUN adduser --disabled-password --gecos "" --home "$(pwd)" \
    --ingroup "users" --no-create-home --uid "888" "nonroot" \
    && chown -R -f 888:888 /app \
    && chmod +x /app/catboostcapi

COPY catboost_model /app/catboost_model

WORKDIR /app
USER 888
CMD ["./catboostcapi"]
