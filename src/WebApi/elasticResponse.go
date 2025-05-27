package webapi

import domain "test-cqrs/src/Domain"

type ElasticResponse struct {
	Hits struct {
		Hits []struct {
			Source domain.Domain `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
