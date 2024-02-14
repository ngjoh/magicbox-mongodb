FROM golang:1.20-buster
WORKDIR /koksmat
COPY . .
WORKDIR /koksmat/.koksmat/app
RUN go install

CMD [ "sleep","infinity"]