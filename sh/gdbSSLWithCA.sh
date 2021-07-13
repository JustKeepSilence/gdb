# ca
openssl genrsa -out ca.key 4096

openssl req -new -sha256 -out ca.csr -key ca.key -config gdb.conf

openssl x509 -req -sha256 -days 365 -in ca.csr -signkey ca.key -out ca.crt -extensions req_ext -extfile gdb.conf

# generate client key file
openssl genrsa -out gdbClient.key 4096

# generate client csr file
openssl req -new -sha256 -out gdbClient.csr -key gdbClient.key -config gdb.conf

# generate client crt file

openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365 -in gdbClient.csr -signkey gdbClient.key -out gdbClient.crt -extensions req_ext -extfile gdb.conf

# generate server key file
openssl genrsa -out gdbServer.key 4096

# generate server csr file
openssl req -new -sha256 -out gdbServer.csr -key gdbServer.key -config gdb.conf

# generate server crt file

openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365 -in gdbServer.csr -signkey gdbServer.key -out gdbServer.crt -extensions req_ext -extfile gdb.conf