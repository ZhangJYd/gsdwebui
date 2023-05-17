package gsdwebui

import (
	"encoding/json"
	"gsdwebui/models"
	"gsdwebui/utils"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func (s *StableDiffusion) Img2Img(req *models.SDImg2imgArgs) (*models.SDResult, error) {
	route, err := url.JoinPath(s.Uri.String(), Img2ImgRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("new img2img task: lens %d %s", len(req.InitImages), req)
	resp, err := s.post(route, req)

	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := &models.SDResult{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, err
}

func NewImg2imgArgs(initImages ...string) *models.SDImg2imgArgs {
	return &models.SDImg2imgArgs{
		InitImages: initImages,
	}
}

func NewControlNetArgs() *models.ControlNetArgs {
	return &models.ControlNetArgs{
		InputImage:    "",
		Mask:          "",
		Module:        "none",
		Model:         "None",
		Weight:        utils.P(1.0),
		ResizeMode:    utils.P(0),
		Lowvram:       utils.P(false),
		ProcessorRes:  utils.P(512),
		ThresholdA:    utils.P(64),
		ThresholdB:    utils.P(64),
		Guidance:      utils.P(1.0),
		GuidanceStart: utils.P(0.0),
		GuidanceEnd:   utils.P(1.0),
		ControlMode:   utils.P(0),
		PixelPerfect:  utils.P(false),
	}
}

func (s *StableDiffusion) Txt2Img(req *models.SDTxt2imgArgs) (*models.SDResult, error) {
	route, err := url.JoinPath(s.Uri.String(), Txt2ImgRoute)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("new txt2img task: %s", req)
	resp, err := s.post(route, req)

	if err != nil || resp.StatusCode != 200 {
		log.Error(route, err, resp.StatusCode, resp.String())
		return nil, err
	}

	ret := &models.SDResult{}
	data := resp.Bytes()
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Error(route, err, string(data))
		return nil, err
	}
	return ret, err
}

func NewTxt2imgArgs() *models.SDTxt2imgArgs {
	return &models.SDTxt2imgArgs{}
}
