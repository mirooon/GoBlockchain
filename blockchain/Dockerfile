FROM golang:1.12-alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN apk add git
RUN go get -d -v github.com/google/uuid
RUN go install -v github.com/google/uuid
RUN go get -d -v github.com/rs/cors
RUN go install -v github.com/rs/cors
RUN go build -o main .
CMD ["./main"]