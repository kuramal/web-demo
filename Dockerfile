#FROM docker.io/centos
FROM docker-registry.default.svc:5000/pro001/test002
#FROM bjdhub.haihangyun.com/openshift/php

ADD ./run /home/run

RUN chmod +x /home/run

EXPOSE 9090

ENTRYPOINT ["/home/run"]
