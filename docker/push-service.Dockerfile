FROM golang:1.17.1

WORKDIR /app

COPY . .

EXPOSE 8080 443 80 5003

RUN make build-push

CMD ./push.out