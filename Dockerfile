FROM golang:1.20-buster
WORKDIR /app.
COPY . .
RUN go install
EXPOSE 5001
CMD [ "koksmat" ,"serve"]