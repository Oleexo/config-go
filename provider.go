package config

type Provider interface {
	Precedence() int
	GetEntry(key string) Entry
}

type Providers []Provider

func (p Providers) Len() int {
	return len(p)
}

func (p Providers) Less(i, j int) bool {
	return p[i].Precedence() < p[j].Precedence()
}
