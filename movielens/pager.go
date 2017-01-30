package movielens

// Pager represents pager object for pagination
type Pager struct {
	ItemsPerPage int `json:"itemsPerPage,omitempty"`
	TotalItems   int `json:"totalItems,omitempty"`
	CurrentPage  int `json:"currentPage,omitempty"`
	TotalPages   int `json:"totalpages,omitempty"`
}
