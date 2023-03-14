package handler

type studentReponse struct {
	ID   int    `json:"id"`
	NPM  int    `json:"npm"`
	Name string `json:"name"`
}

type studentRequest struct {
	NPM  int    `json:"npm,omitempty"`
	Name string `json:"name,omitempty"`
}
