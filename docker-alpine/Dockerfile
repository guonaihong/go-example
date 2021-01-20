FROM alpine:3.7

RUN mkdir -p /opt/service
COPY ./main /opt/service
COPY ./start.sh /opt/service
WORKDIR /opt/service

CMD ["./start.sh"]
