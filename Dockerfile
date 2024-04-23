FROM debian:latest
EXPOSE 3000

RUN mkdir /app
WORKDIR /app

COPY build/linux/amd64/myipserver .

ENTRYPOINT ["/app/myipserver"]