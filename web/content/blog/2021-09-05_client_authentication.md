---
title: "Client authentication"
description: ""
image: "/images/posts/i.imgur.com_zI8WrkR.png"
date: 2021-09-05T08:53:00+07:00
lastmod: 2021-10-03T10:01:00+07:00
author: "Zebra"
tags:
categories:
  - "protocol"
draft: false
---

Client-server encryption and authentication based on Minecraft server.

# Protocol

C: client

S: server

## C → S: Login request

Client requests the login process to start.

## S → C: Encryption request

Server generates a random verify token.

```Go
verifyToken := make([]byte, verifyTokenLength)
rand.Read(verifyToken)
```

Server also generates an unique random 16-byte token for authenticating ownership of client's packets (once the login process has ended). The auth token won't be encrypted, so the server knows the owner of each packet.

Server sends his public key in PKIX, ASN.1 DER format.

```Go
asn1PublicKey, _ := x509.MarshalPXIXPublicKey(global.PublicKey)
```

## C → S: Encryption response

Client generates a shared 16-byte token used as symmetric key, that will be sent to the server encrypted with it's public key. Client also encrypts the verify token.

Once this step is completed, the content of the packet will be encrypted using the symmetric key.

The server decrypts the shared secret and the verify token using its private key:

```Go
secret, _ := rsa.DecryptPKCS1v15(rand.Reader, global.PrivateKey, response.Secret)
verifyToken, _ := rsa.DecryptPKCS1v15(rand.Reader, global.PrivateKey, response.VerifyToken)
```

And then it compares the verify token:

```Go
if bytes.Compare(token, verifyToken) != 0 { ... } // error
```

## C & S: Enable auth

Both enable AES/CFB8 encryption using the shared secret token as both IV and key.

```Go
block, _ := aes.NewCipher(secret)
conn.SetCipher(
	CFB8.NewCFB8Encrypt(block, secret),
	CFB8.NewCFB8Decrypt(block, secret),
)
```

# Questions

- **What if someone intercepts the auth token?** Nothing would happen as the actual authority of packets is determined by the ability to encrypt the packets with the symmetric key.
- **Should authority be considered before login? **No, the encryption ensures the client who requested the connection is the one who is performing it, but unless he logs in, he can't be the user he claims to be.
- **This authenticates the users, but does it authenticate the server? **No, and this should be considered in a future (). If server's public distributio distributed by a trusted entity, clients could be fooled by a malicious server and steal client's password.
