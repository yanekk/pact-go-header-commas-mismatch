package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	. "github.com/pact-foundation/pact-go/v2/consumer"
	. "github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	pact, _ := NewV3Pact(MockHTTPProviderConfig{
		Consumer: "consumer",
		Provider: "producer",
	})

	err := pact.AddInteraction().
		UponReceiving("a request").
		WithRequest("GET", "/endpoint", func(b *V3RequestBuilder) {
			b.Header("Authorization", S("OAuth oauth_consumer_key=abcd, oauth_signature_method=\"PLAINTEXT\", oauth_version=\"1.0\", oauth_signature=\"1234&\""))
		}).
		WillRespondWith(200, func(b *V3ResponseBuilder) {
			b.Body("text/plain", []byte("OK"))
		}).
		ExecuteTest(t, func(mockServerConfig MockServerConfig) error {

			httpClient := &http.Client{}
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/endpoint", mockServerConfig.Host, mockServerConfig.Port), strings.NewReader(""))
			req.Header.Add("Authorization", "OAuth oauth_consumer_key=abcd, oauth_signature_method=\"PLAINTEXT\", oauth_version=\"1.0\", oauth_signature=\"1234&\"")
			_, err := httpClient.Do(req)

			return err
		})

	assert.NoError(t, err)
}
