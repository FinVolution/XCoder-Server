FROM nvidia/cuda:12.2.0-devel-ubuntu20.04
ENV DEBIAN_FRONTEND=noninteractive

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
    sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
    apt-get clean

RUN apt-get update -y && apt-get install -y python3.9 python3.9-distutils curl git vim
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python3.9 get-pip.py
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.9 1

RUN pip3 install vllm