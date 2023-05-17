package models

import (
	"encoding/json"
	"gsdwebui/utils"
)

type SDResult struct {
	Images     []string       `json:"images"`
	Image      string         `json:"image"`
	Parameters map[string]any `json:"parameters"`
	Info       any            `json:"info"`
}

func (s *SDResult) GetImages() []string {
	if s.Images == nil || len(s.Images) != 0 {
		return s.Images
	}
	if len(s.Image) != 0 {
		return []string{s.Image}
	}
	return nil
}

type SDImg2imgArgs struct {
	InitImages             []string `json:"init_images,omitempty"`
	ResizeMode             *int     `json:"resize_mode,omitempty"`
	DenoisingStrength      *float64 `json:"denoising_strength,omitempty"`
	ImageCfgScale          *float64 `json:"image_cfg_scale,omitempty"`
	MaskImage              string   `json:"mask_image,omitempty"`
	MaskBlur               *int     `json:"mask_blur,omitempty"`
	InpaintingFill         *int     `json:"inpainting_fill,omitempty"`
	InpaintFullRes         *bool    `json:"inpaint_full_res,omitempty"`
	InpaintFullResPadding  *int     `json:"inpaint_full_res_padding,omitempty"`
	InpaintingMaskInvert   *int     `json:"inpainting_mask_invert,omitempty"`
	InitialNoiseMultiplier *int     `json:"initial_noise_multiplier,omitempty"`
	IncludeInitImages      *bool    `json:"include_init_images,omitempty"`

	SDCommonArgs
}

func (s *SDImg2imgArgs) AddInitImages(v ...string) *SDImg2imgArgs {
	s.InitImages = append(s.InitImages, v...)
	return s
}

func (s *SDImg2imgArgs) SetResizeMode(v int) *SDImg2imgArgs {
	s.ResizeMode = &v
	return s
}

func (s *SDImg2imgArgs) SetDenoisingStrength(v float64) *SDImg2imgArgs {
	s.DenoisingStrength = &v
	return s
}

func (s *SDImg2imgArgs) SetImageCfgScale(v float64) *SDImg2imgArgs {
	s.ImageCfgScale = &v
	return s
}

func (s *SDImg2imgArgs) SetMaskImage(v string) *SDImg2imgArgs {
	s.MaskImage = v
	return s
}

func (s *SDImg2imgArgs) SetMaskBlur(v int) *SDImg2imgArgs {
	s.MaskBlur = &v
	return s
}

func (s *SDImg2imgArgs) SetInpaintingFill(v int) *SDImg2imgArgs {
	s.InpaintingFill = &v
	return s
}

func (s *SDImg2imgArgs) SetInpaintFullRes(v bool) *SDImg2imgArgs {
	s.InpaintFullRes = &v
	return s
}

func (s *SDImg2imgArgs) SetInpaintFullResPadding(v int) *SDImg2imgArgs {
	s.InpaintFullResPadding = &v
	return s
}

func (s *SDImg2imgArgs) SetInpaintingMaskInvert(v int) *SDImg2imgArgs {
	s.InpaintingMaskInvert = &v
	return s
}

func (s *SDImg2imgArgs) SetInitialNoiseMultiplier(v int) *SDImg2imgArgs {
	s.InitialNoiseMultiplier = &v
	return s
}

func (s *SDImg2imgArgs) SetIncludeInitImages(v bool) *SDImg2imgArgs {
	s.IncludeInitImages = &v
	return s
}

func (s *SDImg2imgArgs) SetPrompt(v string) *SDImg2imgArgs {
	s.Prompt = v
	return s
}

func (s *SDImg2imgArgs) SetNegativePrompt(v string) *SDImg2imgArgs {
	s.NegativePrompt = v
	return s
}

func (s *SDImg2imgArgs) AddStyles(v ...string) *SDImg2imgArgs {
	s.Styles = append(s.Styles, v...)
	return s
}

func (s *SDImg2imgArgs) SetSeed(v int) *SDImg2imgArgs {
	s.Seed = &v
	return s
}

func (s *SDImg2imgArgs) SetSubseed(v int) *SDImg2imgArgs {
	s.Subseed = &v
	return s
}

func (s *SDImg2imgArgs) SetSubseedStrength(v int) *SDImg2imgArgs {
	s.SubseedStrength = &v
	return s
}

func (s *SDImg2imgArgs) SetSeedResizeFromH(v int) *SDImg2imgArgs {
	s.SeedResizeFromH = &v
	return s
}

func (s *SDImg2imgArgs) SetSeedResizeFromW(v int) *SDImg2imgArgs {
	s.SeedResizeFromW = &v
	return s
}

func (s *SDImg2imgArgs) SetSamplerName(v string) *SDImg2imgArgs {
	s.SamplerName = v
	return s
}

func (s *SDImg2imgArgs) SetSamplerIndex(v string) *SDImg2imgArgs {
	s.SamplerIndex = v
	return s
}

func (s *SDImg2imgArgs) SetBatchSize(v int) *SDImg2imgArgs {
	s.BatchSize = &v
	return s
}

func (s *SDImg2imgArgs) SetNIter(v int) *SDImg2imgArgs {
	s.NIter = &v
	return s
}

func (s *SDImg2imgArgs) SetSteps(v int) *SDImg2imgArgs {
	s.Steps = &v
	return s
}

func (s *SDImg2imgArgs) SetCfgScale(v float64) *SDImg2imgArgs {
	s.CfgScale = &v
	return s
}

func (s *SDImg2imgArgs) SetSize(width int, high int) *SDImg2imgArgs {
	s.Height = &high
	s.Width = &width
	return s
}

func (s *SDImg2imgArgs) SetSizeByImgInfo(i *utils.ImageInfo, rate float64) *SDImg2imgArgs {
	s.Height = utils.P(int(float64(i.Height) * rate))
	s.Width = utils.P(int(float64(i.Width) * rate))
	return s
}

func (s *SDImg2imgArgs) SetRestoreFaces(v bool) *SDImg2imgArgs {
	s.RestoreFaces = &v
	return s
}

func (s *SDImg2imgArgs) SetTiling(v bool) *SDImg2imgArgs {
	s.Tiling = &v
	return s
}

func (s *SDImg2imgArgs) SetDoNotSaveSamples(v bool) *SDImg2imgArgs {
	s.DoNotSaveSamples = &v
	return s
}

func (s *SDImg2imgArgs) SetDoNotSaveGrid(v bool) *SDImg2imgArgs {
	s.DoNotSaveGrid = &v
	return s
}

func (s *SDImg2imgArgs) SetEta(v float64) *SDImg2imgArgs {
	s.Eta = &v
	return s
}

func (s *SDImg2imgArgs) SetSChurn(v int) *SDImg2imgArgs {
	s.SChurn = &v
	return s
}

func (s *SDImg2imgArgs) SetSTmax(v int) *SDImg2imgArgs {
	s.STmax = &v
	return s
}

func (s *SDImg2imgArgs) SetSTmin(v int) *SDImg2imgArgs {
	s.STmin = &v
	return s
}

func (s *SDImg2imgArgs) SetSNoise(v int) *SDImg2imgArgs {
	s.SNoise = &v
	return s
}

func (s *SDImg2imgArgs) SetOverrideSettings(k string, v any) *SDImg2imgArgs {
	if s.OverrideSettings == nil {
		s.OverrideSettings = make(map[string]any)
	}
	s.OverrideSettings[k] = v
	return s
}

func (s *SDImg2imgArgs) SetOverrideSettingsRestoreAfterwards(v bool) *SDImg2imgArgs {
	s.OverrideSettingsRestoreAfterwards = &v
	return s
}

func (s *SDImg2imgArgs) SetScriptName(v string) *SDImg2imgArgs {
	s.ScriptName = v
	return s
}

func (s *SDImg2imgArgs) SetScriptArgs(k string, v any) *SDImg2imgArgs {
	if s.ScriptArgs == nil {
		s.ScriptArgs = make(map[string]any)
	}
	s.ScriptArgs[k] = v
	return s
}

func (s *SDImg2imgArgs) SetSendImages(v bool) *SDImg2imgArgs {
	s.SendImages = &v
	return s
}

func (s *SDImg2imgArgs) SetSaveImages(v bool) *SDImg2imgArgs {
	s.SaveImages = &v
	return s
}

func (s *SDImg2imgArgs) SetAlwaysonScripts(k string, v any) *SDImg2imgArgs {
	if s.AlwaysonScripts == nil {
		s.AlwaysonScripts = make(map[string]any)
	}
	s.AlwaysonScripts[k] = v
	return s
}

func (s *SDImg2imgArgs) AddControlNets(cs ...*ControlNetArgs) *SDImg2imgArgs {
	if s.AlwaysonScripts == nil {
		s.AlwaysonScripts = make(map[string]any)
	}

	if len(cs) != 0 {
		if v, ok := s.AlwaysonScripts["ControlNet"].(map[string][]*ControlNetArgs); ok {
			s.AlwaysonScripts["ControlNet"] = map[string][]*ControlNetArgs{
				"args": append(v["args"], cs...),
			}
		} else {
			s.AlwaysonScripts["ControlNet"] = map[string][]*ControlNetArgs{
				"args": cs,
			}
		}
	}
	return s
}

func (s *SDImg2imgArgs) String() string {
	tmp := *s
	tmp.InitImages = nil
	data, _ := json.Marshal(tmp)
	return string(data)
}
