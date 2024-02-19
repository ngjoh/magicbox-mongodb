FROM mcr.microsoft.com/azure-cli
RUN apk add go
RUN apk add powershell
RUN apk add nodejs
RUN apk add npm
WORKDIR /koksmat
COPY . .
WORKDIR /koksmat/.koksmat/app
RUN go install
WORKDIR /koksmat/.koksmat/web
RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build




CMD [ "sleep","infinity"]