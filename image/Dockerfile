FROM docker.io/centos

ADD ./run /root/run 
ADD ./server.crt /root/server.crt
ADD ./server.key /root/server.key

RUN chmod +x /root/run

EXPOSE 443

CMD ["/root/run"]
