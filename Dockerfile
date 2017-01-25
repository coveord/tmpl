FROM alpine:3.5
MAINTAINER Pierre-Alexandre St-Jean <pa@stjean.me>

COPY tmpl /tmpl

ENTRYPOINT ["/tmpl"]