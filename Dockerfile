FROM mcr.microsoft.com/azure-cli
RUN apk add go
RUN apk add powershell

WORKDIR /koksmat
COPY . .
WORKDIR /koksmat/.koksmat/app
RUN go install




CMD [ "sleep","infinity"]