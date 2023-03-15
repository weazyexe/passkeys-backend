# Passkeys Backend demo

Server application for authenticate with Passkeys,
a passwordless way to authenticate using public key cryptography

## How to run

1. Create `.env` file and fill it like `example.env` file.
2. Set your domain in `assets/assetlinks.json`, `generate-certificate.sh` and `assets/self-signed-cert.ext`
3. To run it locally, you should use self-signed
certificates for HTTPS connection. Run:
   ```bash
   ./generate-certificate.sh
   ```
4. If you want to use self-signed certificates, set `true`
to `SELF_SIGNED_CERTIFICATES` variable in `.env`
5. Build the app and run it:
   ```bash
   go build github.com/weazyexe/passkeys/cmd/app && ./app
   ```