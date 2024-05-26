FROM golang:1.22

WORKDIR /Phonebook

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /gorise
CMD ["/gorise"]
