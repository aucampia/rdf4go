# syntax = docker/dockerfile:experimental
# https://docs.docker.com/engine/reference/builder/
# https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md
# https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
ARG base_image_registry=docker.io
ARG base_image_name=alpine
ARG base_image_tag=3
FROM ${base_image_registry}/${base_image_name}:${base_image_tag}

RUN \
    apk --no-cache update &&\
    apk --no-cache add \
        tzdata=2021a-r0 \
    && \
    ln -vfs /usr/share/zoneinfo/UTC /etc/localtime && \
    true

ARG main_cmd

COPY --chown=root:root go/bin/${main_cmd} /usr/local/bin/service

RUN \
    chmod ugo-w /usr/local/bin/service

ARG user=

ARG group_name=eggroup
ARG user_name=eguser
ARG workdir=/var/opt/egsvc

RUN \
    addgroup -S ${group_name} && \
    adduser -S -H -D -h ${workdir} -G ${group_name} ${user_name} && \
    mkdir -vp ${workdir} && \
    chown -R ${user_name}:${group_name} ${workdir} && \
    true

USER ${user_name}:${group_name}
WORKDIR ${workdir}

HEALTHCHECK CMD ["/usr/local/bin/service", "healthcheck"]

ENTRYPOINT ["/usr/local/bin/service"]
CMD ["run"]
