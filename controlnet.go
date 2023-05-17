package gsdwebui

import (
	"encoding/json"
	"gsdwebui/models"
	"net/url"

	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
)

type ControlNet struct {
	StUri *url.URL
}

func newControlNet(stUri ...string) (*ControlNet, error) {
	uri, err := url.Parse("http://127.0.0.1:7860")
	if len(stUri) != 0 {
		uri, err = url.Parse(stUri[0])
	}
	if err != nil {
		return nil, err
	}
	return &ControlNet{
		StUri: uri,
	}, nil
}

func (c *ControlNet) Version() (int, error) {
	route, err := url.JoinPath(c.StUri.String(), ControlNetVersion)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	resp, err := grequests.Get(route, nil)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return 0, err
	}

	var ret map[string]int
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return 0, err
	}
	return ret["version"], nil
}

func (c *ControlNet) ModelList() ([]string, error) {
	route, err := url.JoinPath(c.StUri.String(), ControlNetModelList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := grequests.Get(route, nil)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	var ret models.ControlNetModelListResult
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret.ModelList, nil
}

func (c *ControlNet) ModuleList() ([]string, error) {
	route, err := url.JoinPath(c.StUri.String(), ControlNetModuleList)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := grequests.Get(route, nil)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	var ret models.ControlNetModuleListResult
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret.ModuleList, nil
}

func (c *ControlNet) Detect(req *models.ControlNetDetectArgs) ([]string, error) {
	route, err := url.JoinPath(c.StUri.String(), ControlNetDetect)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resp, err := grequests.Post(route,
		&grequests.RequestOptions{
			JSON: req,
		},
	)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.ControlNetDetectResult{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret.Images, nil
}
