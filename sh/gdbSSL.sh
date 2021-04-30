# generate client key file
openssl genrsa -out gdbClient.key 4096

# generate client csr file
openssl req -new -sha256 -out gdbClient.csr -key gdbClient.key -config gdb.conf

# generate client crt file

openssl x509 -req -sha256  -days 365 -in gdbClient.csr -signkey gdbClient.key -out gdbClient.crt -extensions req_ext -extfile gdb.conf

# generate server key file
openssl genrsa -out gdbServer.key 4096

# generate server csr file
openssl req -new -sha256 -out gdbServer.csr -key gdbServer.key -config gdb.conf

# generate server crt file

openssl x509 -req -sha256 -days 365 -in gdbServer.csr -signkey gdbServer.key -out gdbServer.crt -extensions req_ext -extfile gdb.conf