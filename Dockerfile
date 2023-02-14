FROM golang:latest

WORKDIR /cmd/anti-bruteforce

COPY ./ ./
RUN go build -o ./cmd/anti-bruteforce/app-anti-bruteforce github.com/Srgkharkov/anti-bruteforce/cmd/anti-bruteforce
CMD ["./cmd/anti-bruteforce/main"]