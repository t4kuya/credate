FROM alpine:latest

ARG version=1.0.0

LABEL maintainer="Takuya Yamanaka" \
      description="Creation Date"

RUN    adduser -D credate \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/tamada/credate/realeases/download/v${version}/credate-${version}_linux_amd64.tar.gz \
    && curl -s -L -o credate-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/x1q8swkvd83ycw3/credate-1.0.0_linux_amd64.tar.gz?dl=0 \
    && tar xfz credate-${version}_linux_amd64.tar.gz        \
    && mv credate-${version} /opt                           \
    && ln -s /opt/credate-${version} /opt/credate               \
    && ln -s /opt/credate-${version}/credate /usr/local/bin/credate \
    && rm credate-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/credate"

WORKDIR /home/credate
USER    credate

ENTRYPOINT [ "/opt/credate/credate" ]