package config

type Provider interface {
	Priority() int
	GetEntry(key string) Entry
}

type Providers []Provider

func (p Providers) Len() int {
	return len(p)
}

func (p Providers) Less(i, j int) bool {
	return p[i].Priority() < p[j].Priority()
}
