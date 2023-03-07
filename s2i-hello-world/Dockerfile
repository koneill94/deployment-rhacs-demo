# go-builder
FROM ubi9/go-toolset:1.18

LABEL maintainer="Kieran O'Neill <kieranoneill94@gmail.com>"

ENV GO_VERSION 1.18

LABEL io.k8s.description="Platform for building Go" \
io.k8s.display-name="Helloworld 1.0.0" \
io.openshift.expose-services="8080:http" \
io.openshift.tags="builder,go,Helloworld"

LABEL io.openshift.s2i.scripts-url=image:///usr/libexec/s2i
COPY ./s2i/bin/ /usr/libexec/s2i

RUN chown -R 1001:1001 /opt/app-root

USER 1001

EXPOSE 8080

CMD ["/usr/libexec/s2i/usage"]
