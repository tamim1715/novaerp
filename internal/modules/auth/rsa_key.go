package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

type KeyManager struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewKeyManager initializes RSA keys from PEM strings or generates a new 2048-bit RSA key pair.
func NewKeyManager(privateKeyPEM, publicKeyPEM string) (*KeyManager, error) {
	if privateKeyPEM != "" {
		privKey, err := ParsePrivateKeyFromPEM(privateKeyPEM)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}

		var pubKey *rsa.PublicKey
		if publicKeyPEM != "" {
			pubKey, err = ParsePublicKeyFromPEM(publicKeyPEM)
			if err != nil {
				return nil, fmt.Errorf("failed to parse public key: %w", err)
			}
		} else {
			pubKey = &privKey.PublicKey
		}

		return &KeyManager{
			PrivateKey: privKey,
			PublicKey:  pubKey,
		}, nil
	}

	// Auto-generate RSA 2048 key pair if no PEM is supplied in configuration
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	return &KeyManager{
		PrivateKey: privKey,
		PublicKey:  &privKey.PublicKey,
	}, nil
}

func ParsePrivateKeyFromPEM(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, errors.New("PKCS8 key is not an RSA private key")
	}

	return nil, errors.New("unsupported private key format")
}

func ParsePublicKeyFromPEM(pemStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	if pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes); err == nil {
		return pubKey, nil
	}

	if pub, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		if rsaPubKey, ok := pub.(*rsa.PublicKey); ok {
			return rsaPubKey, nil
		}
		return nil, errors.New("PKIX key is not an RSA public key")
	}

	return nil, errors.New("unsupported public key format")
}
