IP?="192.168.0.1"

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLE=0 go build -o cmd/linux-doh
genssl:
	openssl genrsa -out ${IP}.key 2048
	openssl req -x509 -new -nodes -sha256 -days 3650 -subj "/C=TW/ST=Taiwan/L=Taipei/CN=localhost" -key ${IP}.key -out ${IP}.crt