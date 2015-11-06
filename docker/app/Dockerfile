FROM golang:latest
#FROM golang:1.3-onbuild
MAINTAINER Mario Inga <mario21ic@gmail.com>

# GPM
RUN wget https://raw.githubusercontent.com/pote/gpm/v1.3.2/bin/gpm && chmod +x gpm && mv gpm /usr/local/bin

# GVP
WORKDIR /root
RUN git clone https://github.com/pote/gvp.git
RUN cd /root/gvp && git checkout v0.2.1
RUN cd /root/gvp && ./configure && make install

# Source
WORKDIR /usr/src/myapp
VOLUME /usr/src/myapp
COPY Godeps /usr/src/myapp/Godeps
ENV GOPATH=$GOPATH:/usr/src/myapp:/usr/src/myapp/.godeps
RUN /usr/local/bin/gpm install
#ENTRYPOINT ["/usr/local/bin/gpm install"]
