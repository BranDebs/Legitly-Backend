package json

import (
	"encoding/json"

	"github.com/BranDebs/Legitly-Backend/credibility"
)

type Credibility struct{}

func (c *Credibility) Encode(cred *credibility.Credibility) ([]byte, error) {
	return json.Marshal(cred)
}

func (c *Credibility) Decode(b []byte) (*credibility.Credibility, error) {
	var credibility credibility.Credibility
	if err := json.Unmarshal(b, &credibility); err != nil {
		return nil, err
	}
	return &credibility, nil
}
