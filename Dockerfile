#FROM docker.io/centos
FROM bjdhub.haihangyun.com/openshift/php

ADD ./run /root/run

RUN chmod +x /root/run

EXPOSE 9090

ENTRYPOINT ["/root/run"]
