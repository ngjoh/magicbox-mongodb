FROM golang:1.20-buster

WORKDIR /app.
COPY . .
#RUN go mod download
RUN go build -o .

CMD [ "koksmat" ]