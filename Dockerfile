FROM golang:1.8.3
MAINTAINER Aleksey Salnikov <me@iamsalnikov.ru>

RUN apt-get update &&\
    apt-get install -y zip &&\
    wget https://github.com/yandex/tomita-parser/releases/download/v1.0/libmystem_c_binding.so.linux_x64.zip &&\
    unzip libmystem_c_binding.so.linux_x64.zip &&\
    cp libmystem_c_binding.so /usr/lib/ &&\
    ln -s /usr/lib/libmystem_c_binding.so /usr/lib/libmystem_c_binding.so.1 &&\
    mkdir -p $GOPATH/src/github.com/iamsalnikov/phrase_lemmatiser

COPY ./ $GOPATH/src/github.com/iamsalnikov/phrase_lemmatiser
RUN cd $GOPATH/src/github.com/iamsalnikov/phrase_lemmatiser && \
    go get &&\
    go install github.com/iamsalnikov/phrase_lemmatiser

WORKDIR $GOPATH/src/github.com/iamsalnikov/phrase_lemmatiser

CMD phrase_lemmatiser

EXPOSE 80