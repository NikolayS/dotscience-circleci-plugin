package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dotmesh-io/dotscience-circleci-plugin/pkg/config"
)

// Client is a CircleCI API client to start jobs
type Client interface {
	TriggerNewJob() error
}

// CircleCIV1Client implements several API calls to start build jobs https://circleci.com/docs/api/#trigger-a-new-job
// TODO: implement new preview API https://circleci.com/docs/api/#trigger-a-new-build-by-project-preview
type CircleCIV1Client struct {
	cfg        config.Config
	httpClient *http.Client
}

// New - create new instance of the CircleCI APi client
func New(cfg config.Config, client *http.Client) *CircleCIV1Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &CircleCIV1Client{
		cfg:        cfg,
		httpClient: client,
	}
}

type newJobPayload struct {
	Tag             string
	Revision        string
	BuildParameters map[string]string
}

// TriggerNewJob - triggers new job
func (c *CircleCIV1Client) TriggerNewJob() error {

	payload := newJobPayload{
		Tag:             c.cfg.Tag,
		Revision:        c.cfg.Revision,
		BuildParameters: c.cfg.BuildParameters,
	}

	// if both are set, leaving only tag
	if payload.Tag != "" && payload.Revision != "" {
		payload.Revision = ""
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
