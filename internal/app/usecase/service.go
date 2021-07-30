package usecase

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/EugeneTseitlin/dash-go-code-challenge/internal/app/validation"
	"github.com/EugeneTseitlin/dash-go-code-challenge/internal/p2p/util"
)

// Service is a service which should interact with p2p and / or self-hosted services in a network
type Service struct {
	p2pClient        *http.Client
	selfHostedClient *http.Client
}

// NewService returns a new service
func NewService(p2pClient, selfHostedClient *http.Client) *Service {
	return &Service{
		p2pClient:        p2pClient,
		selfHostedClient: selfHostedClient,
	}
}

// Fetch returns a fetched data
func (s *Service) Fetch(ctx context.Context) ([]map[string]interface{}, error) {

	resp, err := s.p2pClient.Get("http://localhost:8090/data")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decryptionResult, err := decrypt(body)
	if err != nil {
		return nil, err
	}

	var decompressionResult bytes.Buffer
	r, err := gzip.NewReader(bytes.NewReader(decryptionResult))
	if err != nil {
		return nil, err
	}
	io.Copy(&decompressionResult, r)
	r.Close()

	var result []map[string]interface{}
	err = json.Unmarshal(decompressionResult.Bytes(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Store stored a passed items in external service(s)
func (s *Service) Store(ctx context.Context, items []map[string]interface{}) error {

	var err error
	v := validation.CreateValidator()

	err = validation.ValidateData(v, items)
	util.PanicError(err)

	body, err := json.Marshal(items)
	if err != nil {
		return err
	}

	var compressionResult bytes.Buffer
	w := gzip.NewWriter(&compressionResult)
	w.Write(body)
	w.Close()

	encryptionResult, err := encrypt(compressionResult.Bytes())
	if err != nil {
		return err
	}

	_, err = s.p2pClient.Post("http://localhost:8090/data", "application/json", bytes.NewReader(encryptionResult))
	if err != nil {
		return err
	}

	return nil
}

func encrypt(input []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(os.Getenv("AES_SECRET")))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	output := gcm.Seal(nonce, nonce, input, nil)

	return output, nil
}

func decrypt(input []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(os.Getenv("AES_SECRET")))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := input[:nonceSize], input[nonceSize:]
	output, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return output, nil
}
