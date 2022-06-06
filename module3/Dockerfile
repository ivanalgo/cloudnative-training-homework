FROM centos:7

RUN yum update -y 
RUN yum install -y git wget
RUN wget https://golang.google.cn/dl/go1.18.3.linux-amd64.tar.gz && tar -zxvf go1.18.3.linux-amd64.tar.gz && mkdir -p /opt/ && mv go /opt/
ENV PATH /opt/go/bin:$PATH

RUN mkdir -p /opt/source
RUN cd /opt/source && git clone https://github.com/ivanalgo/cloudnative-training-homework.git
WORKDIR /opt/source/cloudnative-training-homework
RUN go build module2/exercise2/main.go

CMD "/opt/source/cloudnative-training-homework/main"
