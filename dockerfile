FROM golang

WORKDIR /app

COPY . .

VOLUME /mnt/shared

RUN go build -o file_server main.go

EXPOSE 8080

CMD [ "./file_server" ]