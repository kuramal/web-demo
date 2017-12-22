FROM docker.io/centos

ADD ./run /root/run

RUN chmod +x /root/run

EXPOSE 9090

ENTRYPOINT ["/root/run"]
~                           
