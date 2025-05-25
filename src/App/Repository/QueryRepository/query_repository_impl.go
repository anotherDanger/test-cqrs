package queryrepository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	helpers "test-cqrs/src/App/Helpers"
	domain "test-cqrs/src/Domain"

	"github.com/sirupsen/logrus"
)

type QueryRepositoryImpl struct{}

func NewQueryRepositoryImpl() QueryRepository {
	return &QueryRepositoryImpl{}
}

func (repo *QueryRepositoryImpl) GetBook(ctx context.Context, params string) ([]*domain.Domain, error) {
	rawJson := fmt.Sprintf(
		`{
			"query":{
				"match":{
					"author": "%s"
				}
			}
		}`, params)
	response, err := http.Post("http://localhost:9200/books/_search", "application/json", bytes.NewBufferString(rawJson))
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/queryrepository", logrus.ErrorLevel, err)
		return nil, err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var results struct {
		Hits struct {
			Hits []struct {
				Source domain.Domain `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/queryrepository", logrus.ErrorLevel, err)
		return nil, err
	}

	var source []*domain.Domain
	for _, hit := range results.Hits.Hits {
		source = append(source, &hit.Source)
	}
	return source, nil

}

func (repo *QueryRepositoryImpl) GetBookByTitle(ctx context.Context, title string) ([]*domain.Domain, error) {
	reqBody := fmt.Sprintf(
		`{
			"query":{
				"bool":{
					"must":[
						{
							"match":{
								"title": "%s"
							}
						}
					]
				}
			}
		}`, title)

	response, err := http.Post("http://localhost:9200/books/_search", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/queryrepository", logrus.ErrorLevel, err)
		return nil, err
	}

	var results struct {
		Hits struct {
			Hits []struct {
				Source domain.Domain `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/queryrepository", logrus.ErrorLevel, err)
		return nil, err
	}
	json.Unmarshal(body, &results)

	var source []*domain.Domain
	for _, hits := range results.Hits.Hits {
		source = append(source, &hits.Source)
	}

	fmt.Println(source)

	return source, nil

}
