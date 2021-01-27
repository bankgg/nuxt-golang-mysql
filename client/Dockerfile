FROM node:15-buster-slim

LABEL maintainer="Kittitach Piyamapornchai <bankgg@gmail.com>"

RUN apt-get update

ENV APP_ROOT /src

WORKDIR ${APP_ROOT}
COPY . ${APP_ROOT}

RUN yarn install
RUN yarn build
RUN yarn cache clean

ENV HOST 0.0.0.0
