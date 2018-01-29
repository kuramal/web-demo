#FROM docker.io/centos
FROM docker-registry.default.svc:5000/pro001/test002
#FROM bjdhub.haihangyun.com/openshift/php

ADD ./run /root/run

RUN chmod +x /root/run

EXPOSE 9090

ENTRYPOINT ["/root/run"]
