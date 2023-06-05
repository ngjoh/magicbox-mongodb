FROM golang:1.20-buster



COPY . .
WORKDIR src
RUN go mod download

RUN go build -o /koksmat

EXPOSE 8080

CMD [ "koksmat" ]