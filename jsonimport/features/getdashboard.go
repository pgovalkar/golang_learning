package features

import (
	"golang_learning/jsonImport/clients/wavefront"
	"golang_learning/jsonImport/parser"

	"github.com/pkg/errors"
)

func ApiRequest(dashboard string, token string, outputpath string) error {

	const (
		defaultWFUrl string = "https://box.wavefront.com/api/v2/"
	)
	c, err := wavefront.New(token, defaultWFUrl)
	if err != nil {
		return errors.WithMessagef(err, "error reading token, please provide wavefront api token")
	}

	respbody, err := c.GetDashboard(dashboard)
	if err != nil {
		return errors.WithMessage(err, "error getting dashboard from wavefront")
	}
	err = parser.YamlParser(respbody.Response, token, outputpath)
	if err != nil {
		return errors.WithMessage(err, "error parsing to yaml file")
	}
	return nil
}
