package params

type ResourceHashMessage struct {
	ContentType string `json:"contentType"`
	Hash        string `json:"ipfsHash"`
	Metadata    string `json:"metadata"`
	Keywords    string `json:"keywords"`
	MessageType string `json:"messageType"`
	Name        string `json:"name"`
	PrivateKey  string `json:"privateKey"`
	PublicKey   string `json:"publicKey"`
}

type Upload struct {
	Metadata   map[string]string `'json:"metadata"`
	Keywords   string            `json:"keywords"`
	PrivateKey string            `json:"senderPrivateKey"`
	PublicKey  string            `json:"receiverPublicKey"`
}
