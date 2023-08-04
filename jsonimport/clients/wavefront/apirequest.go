package wavefront

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	defaultWFUrl string = "https://box.wavefront.com/api/v2/"
)

type Client struct {
	Token   string
	BaseUrl string
}

func New(token, baseurl string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("missing wavefront token")
	}
	if baseurl == "" {
		baseurl = defaultWFUrl
	}
	return &Client{
		Token:   token,
		BaseUrl: baseurl,
	}, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.WithMessagef(err, "failed to do  '%s' request", req.URL.String())
	}
	defer resp.Body.Close()
	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessagef(err, "error reading response body")
	}
	//fmt.Println(string(respdata))

	var respstatus ResponseStatus
	err = json.Unmarshal(respdata, &respstatus)
	if err != nil {
		return nil, errors.WithMessage(err, "Error marshalling status json")
	}
	if respstatus.Status.Result != "OK" {
		return nil, errors.WithMessagef(err, "api request failed to get dashboard with %d : %s", respstatus.Status.Code, respstatus.Status.Message)
	}
	return respdata, nil
}

func (c *Client) get(endpoint string) ([]byte, error) {

	u, err := url.JoinPath(c.BaseUrl, endpoint)
	fmt.Println(u)
	if err != nil {
		return nil, errors.WithMessage(err, "error in url")
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to connet to wavefront endpoint %s", u)
	}

	return c.do(req)
}

func (c *Client) post(endpoint string, body []byte) ([]byte, error) {

	u, err := url.JoinPath(c.BaseUrl, endpoint)
	//fmt.Println(u)
	if err != nil {
		return nil, errors.WithMessage(err, "error in url")
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to connet to wavefront endpoint %s", u)
	}
	//fmt.Println(req)
	return c.do(req)
}

func (c *Client) GetDashboard(id string) (*DashboardResponse, error) {
	endpoint := "dashboard/" + id

	resbody, err := c.get(endpoint)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to connect o wavefront endpoint %s ", endpoint)
	}
	//fmt.Println(string(resbody))
	var dashboard DashboardResponse
	//var dashboard WaveFrontDashboardSpec
	if err := json.Unmarshal(resbody, &dashboard); err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal json response body")
	}

	// //Marshal the file
	// output, err := json.MarshalIndent(dashboard, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// }
	// //writing sub-section key/values to file
	// err = core.WriteFile(dashboard.Response.Url, output, "json_temp")
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// }

	return &dashboard, nil

}

// TODO
func (c *Client) GetAlert(tags string) error {
	endpoint := "search/alert"
	body := []byte(fmt.Sprintf(`{"limit":1000,
	"query":[{"key":"tags","value":"%s","matchingMethod":"CONTAINS","negated":false}],
	"sort":{"ascending":true,"field":"default"}}`, tags))

	resbody, err := c.post(endpoint, body)
	if err != nil {
		return errors.WithMessagef(err, "failed to connect o wavefront endpoint %s ", endpoint)
	}

	//fmt.Println(string(resbody))
	var alert AlertResponse

	if err := json.Unmarshal(resbody, &alert); err != nil {
		return errors.WithMessage(err, "failed to unmarshal json response body")
	}

	fmt.Println(alert.Response.Items)
	//Marshal the file
	output, err := json.MarshalIndent(alert, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println(string(output))
	//writing sub-section key/values to file
	// err = core.WriteFile(alert.Response, output, "json_temp")
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// }
	return nil
}
