FROM golang:1.16.3-buster

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go get github.com/cosmtrek/air
RUN go get -u github.com/labstack/echo
RUN go get github.com/jinzhu/gorm
RUN go get github.com/rubenv/sql-migrate/...
CMD ["air"]