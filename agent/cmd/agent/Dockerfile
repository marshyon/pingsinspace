FROM golang
MAINTAINER marshyon
LABEL Remarks="This is a dockerfile example for an Ubuntu system"
RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y monitoring-plugins && \
    apt-get install -y iputils-ping && \
    apt-get install -y net-tools && \
    apt-get install -y uuid-runtime && \
    mkdir /root/agent
COPY main.go /root/agent/main.go
COPY jobs.json /root/agent/jobs.json
RUN cd /root/agent && go build
EXPOSE 80
ENV HOME /root
WORKDIR /root
#ENTRYPOINT ["ping"]
#CMD ["yahoo.com"]
