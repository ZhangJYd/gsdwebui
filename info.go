package gsdwebui

import (
	"encoding/json"
	"gsdwebui/models"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func (s *StableDiffusion) GetOptions() (*models.StableDiffusionOptions, error) {
	route, err := url.JoinPath(s.Uri.String(), OptionsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.StableDiffusionOptions{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}

	return &ret, nil
}

func (s *StableDiffusion) GetOptions2() (map[string]any, error) {
	route, err := url.JoinPath(s.Uri.String(), OptionsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make(map[string]any)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}

	return ret, nil
}

func (s *StableDiffusion) GetSDModels() ([]models.SDModelInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), SDModelsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.SDModelInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetScripts() (*models.ScriptsInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), ScriptsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.ScriptsInfo{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return &ret, nil
}

func (s *StableDiffusion) GetHypenetworks() ([]models.HypernetworksInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), HypernetRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.HypernetworksInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetFaceRestores() ([]models.FaceRestoresInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), FaceRestoresRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.FaceRestoresInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetSamplers() ([]models.SamplerInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), SamplersRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.SamplerInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetProgress() (*models.ProgressInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), ProgressRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.ProgressInfo{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return &ret, nil
}

func (s *StableDiffusion) GetCmdFlags() (*models.CmdFlagsInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), CmdFlagsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.CmdFlagsInfo{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return &ret, nil
}

func (s *StableDiffusion) GetUpscalers() ([]models.UpscalerInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), UpscalersRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.UpscalerInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetRealesrganModels() ([]models.RealesrganModelInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), RealesrganModelsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.RealesrganModelInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetPromptStyles() ([]models.PromptStyleInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), PromptStylesRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.PromptStyleInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}

func (s *StableDiffusion) GetEmbeddings() (*models.EmbeddingsInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), EmbeddingsRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.EmbeddingsInfo{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return &ret, nil
}

func (s *StableDiffusion) GetMemoryInfo() (*models.MemoryInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), MemoryRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := models.MemoryInfo{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return &ret, nil
}

func (s *StableDiffusion) GetLoras() ([]models.LoraInfo, error) {
	route, err := url.JoinPath(s.Uri.String(), LoraRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := s.get(route)
	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := make([]models.LoraInfo, 0)
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, nil
}
