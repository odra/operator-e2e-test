ARG BASEIMAGE

FROM ${BASEIMAGE}

ADD tmp/_output/bin/stuff-test /usr/local/bin/stuff-test
ARG NAMESPACEDMAN
ADD $NAMESPACEDMAN /namespaced.yaml
ADD tmp/build/go-test.sh /go-test.sh
