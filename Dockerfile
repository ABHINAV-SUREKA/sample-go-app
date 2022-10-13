FROM golang:1.16-alpine
WORKDIR /app
COPY *.go ./
RUN go build -o /tsdb_api
EXPOSE 8080
CMD [ "/tsdb_api" ]
