package dtos

type CreateMovieRequest struct {
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	Description string  `json:"description"`
	Poster      string  `json:"poster"`
	Trailer     string  `json:"trailer"`
	Duration    int     `json:"duration"`
	Rating      float64 `json:"rating"`
	FunFacts    string  `json:"fun_facts"`
	Grade       string  `json:"grade"`
}

type MovieResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	Description string  `json:"description"`
	Poster      string  `json:"poster"`
	Trailer     string  `json:"trailer"`
	Duration    int     `json:"duration"`
	Rating      float64 `json:"rating"`
	FunFacts    string  `json:"fun_facts"`
	Grade       string  `json:"grade"`
}
