package store

import (
	"net/http"
	"strconv"
)

type PaginationFeedQuery struct {
	Limit  int `json:"limit" validate:"gte=1,lte=20"`
	Offset int `json:"offset" validate:"gte=0"`
	Sort string `json:"sort" validate:"oneof=asc desc"`
}

func (p PaginationFeedQuery) Parse(r *http.Request) (PaginationFeedQuery, error) { 
	queryString := r.URL.Query()
	
	limit := queryString.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return p, err
		}
		p.Limit = l
	}

	offset := queryString.Get("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return p, err
		}
		p.Offset = o
	}

	sort := queryString.Get("sort")
	if sort != "" {
		p.Sort = sort
	}

	return p, nil
}