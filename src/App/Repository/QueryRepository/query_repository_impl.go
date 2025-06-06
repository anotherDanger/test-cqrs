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
	webapi "test-cqrs/src/WebApi"

	"github.com/sirupsen/logrus"
)

type QueryRepositoryImpl struct{}

func NewQueryRepositoryImpl() QueryRepository {
	return &QueryRepositoryImpl{}
}

func (repo *QueryRepositoryImpl) GetBook(ctx context.Context, key string, value string) ([]*domain.Domain, error) {
	rawJson := fmt.Sprintf(
		`{
			"query":{
				"match":{
					"%s": "%s"
				}
			}
		}`, key, value)
	response, err := http.Post("http://host.docker.internal:9200/books/_search", "application/json", bytes.NewBufferString(rawJson))
	if err != nil {
		helpers.NewErr("/app/src/App/logs/repository", logrus.ErrorLevel, err)
		return nil, err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var results webapi.ElasticResponse

	err = json.Unmarshal(body, &results)
	if err != nil {
		helpers.NewErr("/app/src/App/logs/repository", logrus.ErrorLevel, err)
		return nil, err
	}

	var source []*domain.Domain
	for _, hit := range results.Hits.Hits {
		source = append(source, &hit.Source)
	}
	return source, nil

}
