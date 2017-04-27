openssl genrsa -out ssl.key 2048
openssl req -new -key ssl.key -out ssl.csr -subj "/CN=localhost"
openssl x509 -req -days 3650 -in ssl.csr -signkey ssl.key -out ssl.crt