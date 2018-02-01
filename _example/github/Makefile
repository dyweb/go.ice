.PHONY: fmt
fmt:
	gofmt -d -l -w .

.PHONY: keygen
keygen:
	openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout icehub.key -out icehub.crt -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=someone@random.com"
