package credibility

type credibilityService struct {
	credRepo Repository
}

func NewCredibilityService(repo Repository) CredibilityService {
	return &credibilityService{
		credRepo: repo,
	}
}

func (svc *credibilityService) Find(URL string) (*Credibility, error) {
	return svc.credRepo.Get(URL)
}

func (svc *credibilityService) Store(cred *Credibility) error {
	return svc.credRepo.Add(cred)
}
