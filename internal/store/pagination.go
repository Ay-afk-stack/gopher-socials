package store

import (
	"net/http"
	"strconv"
	"strings"
)

type PaginationFeedQuery struct {
	Limit  int `json:"limit" validate:"gte=1,lte=20"`
	Offset int `json:"offset" validate:"gte=0"`
	Sort string `json:"sort" validate:"oneof=asc desc"`
	Tags []string `json:"tags" validate:"max=5"`
	Search string `json:"search" validate:"max=100"`
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

	tags := queryString.Get("tags")
	if tags != "" {
		p.Tags = strings.Split(tags, ",")
	}

	search := queryString.Get("search")
	if search != "" {
		p.Search = search
	}

	return p, nil
}
