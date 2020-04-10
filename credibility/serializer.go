package credbility

type CredibilitySerializer interface {
	Encode(*Credibility) ([]byte, error)
	Decode([]byte]) (*Credibility, error)
}
