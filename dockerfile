FROM golang:1.26
EXPOSE 6767
WORKDIR /url-shortener
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build
CMD ["make", "run-server"]
