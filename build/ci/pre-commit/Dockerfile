FROM ubuntu

ENV TZ=Europe/London
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    git \
    libxml2-dev \
    libxslt1-dev \
    nodejs \
    ruby \
    ruby-dev \
    unzip \
    xz-utils \
    python3.6 \
    python3-pip \
    wget \
    zlib1g-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt

RUN wget https://dl.google.com/go/go1.15.3.linux-amd64.tar.gz
RUN tar -xvf go1.15.3.linux-amd64.tar.gz
RUN mv go /usr/local

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/.go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN go get github.com/fzipp/gocyclo/cmd/gocyclo
RUN GO111MODULE=on go get -v -u github.com/go-critic/go-critic/cmd/gocritic


RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0

RUN pip3 install pre-commit==2.9.3

WORKDIR /pre-commit

ENTRYPOINT ["pre-commit"]
