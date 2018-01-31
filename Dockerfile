#FROM docker.io/centos
FROM docker-registry.default.svc:5000/pro001/test002
#FROM bjdhub.haihangyun.com/openshift/php

ADD ./run /opt/app-root/src/run

#RUN chmod +x /opt/app-root/src/run

EXPOSE 9090

ENTRYPOINT ["/opt/app-root/src/run"]
