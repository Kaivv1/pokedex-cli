package pokedex

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokedex: make(map[string]Pokemon),
	}
}

func (p *Pokedex) Add(pokemon Pokemon) {
	p.Pokedex[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(pokemonName string) (Pokemon, bool) {
	pokemon, exists := p.Pokedex[pokemonName]
	return pokemon, exists
}

func (p *Pokedex) List() map[string]Pokemon {
	return p.Pokedex
}
