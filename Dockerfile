FROM golang:1.23-alpine
WORKDIR /wallet-api
COPY . /wallet-api
RUN go build /wallet-api
EXPOSE 8080
ENTRYPOINT [ "./wallet-api" ]