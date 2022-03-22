FROM node:lts-alpine3.11

RUN apk add bind-tools 
RUN apk add --no-cache nodejs npm
RUN mkdir -p /app
WORKDIR /app

# copy source files
COPY ./package*.json /app/

# install dependencies
RUN npm install

COPY . /app/
# start app
EXPOSE 3000
CMD npm run dev