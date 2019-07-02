FROM alpine:3.10.0 AS Runner

WORKDIR /app
RUN apk add --update --no-cache tzdata \
 && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
 && echo "Asia/Tokyo" > /etc/timezone \
    apk del tzdata \
 && apk add --no-cache sqlite sqlite-dev \
 && touch /app/thanks.db
COPY ./thanks-server /app/thanks-server

CMD /app/thanks-server
