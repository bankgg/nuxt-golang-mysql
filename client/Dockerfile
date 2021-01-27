FROM node:14.11-buster-slim

LABEL maintainer="Kittitach Piyamapornchai <bankgg@gmail.com>"

RUN apt-get update

ENV APP_ROOT /src

WORKDIR ${APP_ROOT}
COPY . ${APP_ROOT}

RUN yarn build

ENV HOST 0.0.0.0
