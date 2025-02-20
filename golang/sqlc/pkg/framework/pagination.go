package framework

type PaginatedRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p PaginatedRequest) GetPage() int {
	return p.Page - 1
}

func (p PaginatedRequest) GetLimit() int {
	return p.Limit
}

func (p PaginatedRequest) GetOffset() int {
	if p.Page > 0 {
		return (p.Page - 1) * p.Limit
	}
	return p.Page
}
