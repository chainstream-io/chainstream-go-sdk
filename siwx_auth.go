package chainstream

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	siwxDomain           = "api.chainstream.io"
	siwxURI              = "https://api.chainstream.io"
	siwxVersion          = "1"
	defaultChainID       = "8453"
	defaultExpiresIn     = time.Hour
	refreshBeforeExpiry  = 60 * time.Second
)

// BuildSiwxMessage constructs an EIP-4361 (SIWE/SIWS) formatted message.
func BuildSiwxMessage(address, chain, nonce string, issuedAt, expiresAt time.Time) string {
	chainLabel := "Ethereum"
	chainID := defaultChainID
	if chain == "solana" {
		chainLabel = "Solana"
		chainID = "mainnet"
	}

	return fmt.Sprintf(`%s wants you to sign in with your %s account:
%s

Sign in to ChainStream API

URI: %s
Version: %s
Chain ID: %s
Nonce: %s
Issued At: %s
Expiration Time: %s`,
		siwxDomain, chainLabel, address,
		siwxURI, siwxVersion, chainID, nonce,
		issuedAt.UTC().Format(time.RFC3339),
		expiresAt.UTC().Format(time.RFC3339))
}

// CreateSiwxToken generates a SIWX token: base64(message).signature
func CreateSiwxToken(signer WalletSigner, expiresIn time.Duration) (string, time.Time, error) {
	nonce := generateNonce()
	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(expiresIn)

	message := BuildSiwxMessage(signer.Address(), signer.Chain(), nonce, issuedAt, expiresAt)
	signature, err := signer.SignMessage(message)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("SIWX sign failed: %w", err)
	}

	messageB64 := base64.StdEncoding.EncodeToString([]byte(message))
	token := messageB64 + "." + signature
	return token, expiresAt, nil
}

// SiwxTokenCache caches SIWX tokens with automatic refresh.
type SiwxTokenCache struct {
	mu        sync.Mutex
	token     string
	expiresAt time.Time
	signer    WalletSigner
}

// NewSiwxTokenCache creates a cache for the given signer.
func NewSiwxTokenCache(signer WalletSigner) *SiwxTokenCache {
	return &SiwxTokenCache{signer: signer}
}

// GetToken returns a valid SIWX token, refreshing if needed.
func (c *SiwxTokenCache) GetToken() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != "" && time.Now().Add(refreshBeforeExpiry).Before(c.expiresAt) {
		return c.token, nil
	}

	token, expiresAt, err := CreateSiwxToken(c.signer, defaultExpiresIn)
	if err != nil {
		return "", err
	}
	c.token = token
	c.expiresAt = expiresAt
	return token, nil
}

func siwxGenerateNonce() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// SiwxAuthHeaderFn returns a request editor that adds Authorization: SIWX header.
func SiwxAuthHeaderFn[T ~func(context.Context, *http.Request) error](cache *SiwxTokenCache) T {
	return T(func(_ context.Context, req *http.Request) error {
		token, err := cache.GetToken()
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("SIWX %s", token))
		req.Header.Set("Content-Type", "application/json")
		return nil
	})
}
