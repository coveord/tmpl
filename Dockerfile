FROM alpine:3.5
MAINTAINER Pierre-Alexandre St-Jean <pa@stjean.me>

COPY tmpl.linux.amd64 /tmpl

ENTRYPOINT ["/tmpl"]