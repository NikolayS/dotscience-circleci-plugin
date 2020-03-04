package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dotmesh-io/dotscience-circleci-plugin/pkg/config"
)

type Client interface {
	TriggerNewJob()
}

type TriggerNewJobOpts struct {
	Username string
	Token    string
	Project  string
	VCSType  string

	Revision        string            // optional
	Tag             string            // optional
	BuildParameters map[string]string // optional
}

// CircleCIV1Client implements several API calls to start build jobs https://circleci.com/docs/api/#trigger-a-new-job
// TODO: implement new preview API https://circleci.com/docs/api/#trigger-a-new-build-by-project-preview
type CircleCIV1Client struct {
	cfg        config.Config
	httpClient *http.Client
}

type newJobPayload struct {
	Tag             string
	Revision        string
	BuildParameters map[string]string
}

// TriggerNewJob - triggers new job
func (c *CircleCIV1Client) TriggerNewJob() error {

	payload := newJobPayload{
		Tag:             c.cfg.Token,
		Revision:        c.cfg.Revision,
		BuildParameters: c.cfg.BuildParameters,
	}

	bts, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	var req *http.Request

	if c.cfg.Branch != "" {
		req, err = http.NewRequest(http.MethodPost,
			// https://circleci.com/api/v1.1/project/:vcs-type/:username/:project/tree/:branch?circle-token=:token
			fmt.Sprintf("%s/api/v1.1/project/%s/%s/%s/tree/%s?circle-token=%s", c.cfg.Host, c.cfg.VCSType, c.cfg.Username, c.cfg.Project, c.cfg.Branch, c.cfg.Token),
			bytes.NewBuffer(bts),
		)
	} else {
		req, err = http.NewRequest(http.MethodPost,
			// https://circleci.com/api/v1.1/project/:vcs-type/:username/:project?circle-token=:token
			fmt.Sprintf("%s/api/v1.1/project/%s/%s/%s?circle-token=%s", c.cfg.Host, c.cfg.VCSType, c.cfg.Username, c.cfg.Project, c.cfg.Token),
			bytes.NewBuffer(bts),
		)
	}

	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call CircleCI API (%s), error: %s", c.cfg.Host, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		respBodyBts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected status code %d (failed to read response body)", resp.StatusCode)
		}
		return fmt.Errorf("unexpected status code %d (%s)", resp.StatusCode, string(respBodyBts))
	}

	return nil
}
