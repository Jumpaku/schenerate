FROM golang:1.21.5-bullseye AS work-base

ENV DEBIAN_FRONTEND=noninteractive

CMD ["bash"]

WORKDIR /work

RUN apt update && apt install -y git curl jq make sqlite3 python3 postgresql clang

