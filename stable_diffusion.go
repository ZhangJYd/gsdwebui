package gsdwebui

import (
	"fmt"
	"gsdwebui/models"
	"strings"
	"sync"
	"time"

	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"

	"net/url"
)

type StableDiffusion struct {
	Uri        *url.URL
	ControlNet *ControlNet
	xm         *sync.RWMutex
}

func NewStableDiffusion(stUri ...string) (*StableDiffusion, error) {
	uri, err := url.Parse("http://127.0.0.1:7860")
	if len(stUri) != 0 {
		uri, err = url.Parse(stUri[0])
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	ca, err := newControlNet(stUri...)
	if err != nil {
		log.Error()
		return nil, err
	}
	return &StableDiffusion{
		Uri:        uri,
		ControlNet: ca,
		xm:         &sync.RWMutex{},
	}, nil
}

func (s *StableDiffusion) ControlNetVersion() (int, error) {
	s.xm.RLock()
	defer s.xm.RUnlock()
	return s.ControlNet.Version()
}

func (s *StableDiffusion) ControlNetModelList() ([]string, error) {
	s.xm.RLock()
	defer s.xm.RUnlock()
	return s.ControlNet.ModelList()
}

func (s *StableDiffusion) FindControlNetModels(k string) ([]string, error) {
	models, err := s.ControlNetModelList()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(models))

	for _, model := range models {
		if strings.Contains(model, k) {
			ret = append(ret, model)
		}
	}
	if len(ret) == 0 {
		ret = append(ret, "None")
		return ret, fmt.Errorf("ControlNet model not found")
	}

	return ret, nil
}

func (s *StableDiffusion) FindFirstControlNetModel(k string) (string, error) {
	models, err := s.ControlNetModelList()
	if err != nil {
		return "", err
	}

	for _, model := range models {
		if strings.Contains(model, k) {
			return model, nil
		}
	}
	return "None", fmt.Errorf("ControlNet model not found")
}

func (s *StableDiffusion) ControlNetModuleList() ([]string, error) {
	s.xm.RLock()
	defer s.xm.RUnlock()
	return s.ControlNet.ModuleList()
}

func (s *StableDiffusion) ControlNetDetect(req *models.ControlNetDetectArgs) ([]string, error) {
	s.xm.RLock()
	defer s.xm.RUnlock()
	return s.ControlNet.Detect(req)
}

func (s *StableDiffusion) get(url string) (*grequests.Response, error) {
	s.xm.RLock()
	defer s.xm.RUnlock()
	return grequests.Get(url, nil)
}

func (s *StableDiffusion) post(url string, JsonPayload any) (*grequests.Response, error) {
	s.xm.Lock()
	defer s.xm.Unlock()
	return grequests.Post(url, &grequests.RequestOptions{
		JSON:        JsonPayload,
		DialTimeout: time.Hour * 24,
	})
}

func (s *StableDiffusion) SetOptions(req *models.StableDiffusionOptions) error {
	route, err := url.JoinPath(s.Uri.String(), OptionsRoute)
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := s.post(route, req)

	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return err
	}
	log.Infof("set options %v success: %s", req, resp.String())
	return nil
}

func (s *StableDiffusion) SetOptions2(req map[string]any) error {
	route, err := url.JoinPath(s.Uri.String(), OptionsRoute)
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := s.post(route, req)

	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return err
	}
	log.Infof("set options %v success: %s", req, resp.String())
	return nil
}

func (s *StableDiffusion) SetOptionsAny(req map[string]any) error {
	route, err := url.JoinPath(s.Uri.String(), OptionsRoute)
	if err != nil {
		log.Error(err)
		return err
	}

	resp, err := s.post(route, req)

	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return err
	}
	log.Infof("set options %v success: %s", req, resp.String())
	return nil
}

func (s *StableDiffusion) SetOption(option string, values any) error {
	return s.SetOptionsAny(map[string]any{option: values})
}

func (s *StableDiffusion) SetSDModelCheckpoint(req *models.SDModelInfo) error {
	return s.SetOptions(&models.StableDiffusionOptions{
		SdModelCheckpoint: req.Hash,
	})
}
