FROM debian:latest
EXPOSE 3000

RUN mkdir /app
WORKDIR /app

COPY build/linux/amd64/myipserver .

RUN mkdir /data
COPY data /data

ENTRYPOINT ["/app/myipserver"]