package chainstream

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const signaturePrefix = "chainstream"

// WalletSigner signs messages for x402 wallet signature authentication.
type WalletSigner interface {
	// Chain returns "evm" or "solana".
	Chain() string
	// Address returns the wallet address (hex for EVM, Base58 for Solana).
	Address() string
	// SignMessage signs the message and returns the signature.
	// EVM: personal_sign (EIP-191), return 0x-prefixed hex.
	// Solana: ed25519 signMessage, return hex or base58.
	SignMessage(message string) (string, error)
}

// BuildSignMessage constructs the message that the agent should sign.
func BuildSignMessage(chain, address, timestamp, nonce string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", signaturePrefix, chain, address, timestamp, nonce)
}

func generateNonce() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// WalletAuthHeaderFn returns a request editor function that adds X-Wallet-* headers.
// Each invocation generates a fresh timestamp and nonce, so it is safe for repeated use.
func WalletAuthHeaderFn[T ~func(context.Context, *http.Request) error](signer WalletSigner) T {
	return T(func(_ context.Context, req *http.Request) error {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		nonce := generateNonce()
		message := BuildSignMessage(signer.Chain(), signer.Address(), timestamp, nonce)

		signature, err := signer.SignMessage(message)
		if err != nil {
			return fmt.Errorf("wallet sign failed: %w", err)
		}

		req.Header.Set("X-Wallet-Address", signer.Address())
		req.Header.Set("X-Wallet-Chain", signer.Chain())
		req.Header.Set("X-Wallet-Signature", signature)
		req.Header.Set("X-Wallet-Timestamp", timestamp)
		req.Header.Set("X-Wallet-Nonce", nonce)
		req.Header.Set("Content-Type", "application/json")
		return nil
	})
}
