FILE_CERT_NAME=passkeys_cert
CERT_DOMAIN=surface.me

if [ -f "assets/$FILE_CERT_NAME.crt" ] && [ -f "assets/$FILE_CERT_NAME.key" ] ; then
    echo "Cert and Key already exist"
else
    echo "Cert and Key does not exist, trying to create new ones..."
    openssl req -new -subj "/C=GE/CN=$CERT_DOMAIN" \
        -newkey rsa:2048 -nodes -keyout "assets/$FILE_CERT_NAME.key" -out "assets/$FILE_CERT_NAME.csr"
    openssl x509 -req -days 365 -in "assets/$FILE_CERT_NAME.csr" -signkey "assets/$FILE_CERT_NAME.key" -out "assets/$FILE_CERT_NAME.crt" -extfile "assets/self-signed-cert.ext"
fi