FROM alpine:latest

RUN mkdir /app

COPY pismo /app

CMD [ "/app/pismo" ]