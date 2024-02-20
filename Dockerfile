FROM mcr.microsoft.com/azure-cli
# install the requirements
RUN apk update
RUN apk add --upgrade powershell   
# Start PowerShell


#RUN apk add nodejs
#RUN apk add nodejs
#RUN apk add npm
WORKDIR /koksmat
COPY . .
WORKDIR /koksmat/.koksmat/app
#RUN go install
WORKDIR /koksmat/.koksmat/web
#RUN npm install -g pnpm
#RUN pnpm install
#RUN pnpm build




CMD [ "sleep","infinity"]