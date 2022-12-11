FROM alpine:3.17.0
WORKDIR /apps
COPY bin/cryprom /apps/cryprom
EXPOSE 2112
ENV COINS="bitcoin,ethereum"
ENTRYPOINT ["/apps/cryprom"]