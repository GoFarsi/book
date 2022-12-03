FROM nginx:alpine

# Labels
LABEL maintainer="Ja7ad <ja7ad@live.com>"
LABEL description="gofarsi book offline in docker"
LABEL org.opencontainers.image.source https://github.com/GoFarsi/book

RUN apk add --no-cache git
RUN git clone https://github.com/GoFarsi/book && cd book && git checkout gh-pages
WORKDIR /book
RUN cp -r . /usr/share/nginx/html
