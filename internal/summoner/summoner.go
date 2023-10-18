package summoner

type Summoner struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	PuuID string `json:"puuid"`
}

func (s *Summoner) UpdateName(name string) {
	s.Name = name
}

func (s *Summoner) PullName() string {
	return s.Name
}

func (s *Summoner) PullID() string {
	return s.ID
}

func (s *Summoner) PullPUUID() string {
	return s.PuuID
}
