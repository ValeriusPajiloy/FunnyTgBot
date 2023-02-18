FROM golang:1.19

RUN go version

RUN apt-get update
RUN apt install glibc-source -y

WORKDIR /app

COPY bin .

CMD [ "./main" ]
