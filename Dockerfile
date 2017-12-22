FROM docker.io/centos

ADD ./run /root/run

RUN chmod +x /root/run

ENTRYPOINT ["/root/run"]
~                           
