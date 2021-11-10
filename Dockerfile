FROM scratch
LABEL Maintainer="muthu <gkmuthukumaran@gmail.com>"
WORKDIR /app
COPY ./blogapi /
COPY ./config.yml /
ENV PORT=8080
EXPOSE $PORT
CMD ["/blogapi"]