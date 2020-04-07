package credibility

type CredibilityService interface {
	Find(URL string) (*Credibility, error)
	Store(*Credibility) error
}
