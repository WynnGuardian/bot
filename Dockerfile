FROM golang:1.23.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY ./bot /go/src

EXPOSE 8087

CMD ["tail", "-f", "/dev/null"]