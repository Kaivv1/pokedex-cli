package types

type ClientCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type Config struct {
	Next     *string
	Previous *string
}

type LocationArea struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []Area  `json:"results"`
}

type Area struct {
	Name string `json:"name"`
}
