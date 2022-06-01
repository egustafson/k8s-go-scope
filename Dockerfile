FROM golang:buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /k8s-go-scope

EXPOSE 8080

CMD [ "/k8s-go-scope" ]
