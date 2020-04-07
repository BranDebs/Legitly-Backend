package credibility

type Repository interface {
	Add(*Credibility) error
	Get(URL string) (*Credibility, error)
	Update(*Credibility) (*Credibility, error)
}
