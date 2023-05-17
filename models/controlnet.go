package models

import "gsdwebui/utils"

type ControlNetDetectArgs struct {
	ControlNetModule       string   `json:"controlnet_module,omitempty"`
	ControlNetInputImages  []string `json:"controlnet_input_images,omitempty"`
	ControlNetProcessorRes *int     `json:"controlnet_processor_res,omitempty"`
	ControlNetThresholdA   *int     `json:"controlnet_threshold_a,omitempty"`
	ControlNetThresholdB   *int     `json:"controlnet_threshold_b,omitempty"`
}

func DefaultControlNetDetectArgs() *ControlNetDetectArgs {
	return &ControlNetDetectArgs{
		ControlNetModule:       "",
		ControlNetInputImages:  nil,
		ControlNetProcessorRes: utils.P(512),
		ControlNetThresholdA:   utils.P(64),
		ControlNetThresholdB:   utils.P(64),
	}
}

type ControlNetDetectResult struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}

type ControlNetModuleListResult struct {
	ModuleList   []string    `json:"module_list"`
	ModuleDetail interface{} `json:"module_detail"`
}

type ControlNetModelListResult struct {
	ModelList []string `json:"model_list"`
}

type ControlNetArgs struct {
	InputImage string   `json:"input_image,omitempty"`
	Mask       string   `json:"mask,omitempty"`
	Module     string   `json:"module,omitempty"`
	Model      string   `json:"model,omitempty"`
	Weight     *float64 `json:"weight,omitempty"`
	// 0 Just resize; 1 Scale to Fit; 2 Envelope
	ResizeMode    *int     `json:"resize_mode,omitempty"`
	Lowvram       *bool    `json:"lowvram,omitempty"`
	ProcessorRes  *int     `json:"processor_res,omitempty"`
	ThresholdA    *int     `json:"threshold_a,omitempty"`
	ThresholdB    *int     `json:"threshold_b,omitempty"`
	Guidance      *float64 `json:"guidance,omitempty"`
	GuidanceStart *float64 `json:"guidance_start,omitempty"`
	GuidanceEnd   *float64 `json:"guidance_end,omitempty"`
	// 0 Balanced; 1 My prompt is more important; 2 ControlNet is more important
	ControlMode  *int  `json:"control_mode,omitempty"`
	PixelPerfect *bool `json:"pixel_perfect,omitempty"`
}

func (c *ControlNetArgs) SetInputImage(v string) *ControlNetArgs {
	c.InputImage = v
	return c
}

func (c *ControlNetArgs) SetMask(v string) *ControlNetArgs {
	c.Mask = v
	return c
}

func (c *ControlNetArgs) SetModule(v string) *ControlNetArgs {
	c.Module = v
	return c
}

func (c *ControlNetArgs) SetModel(v string) *ControlNetArgs {
	c.Model = v
	return c
}

func (c *ControlNetArgs) SetWeight(v float64) *ControlNetArgs {
	c.Weight = &v
	return c
}

func (c *ControlNetArgs) SetResizeMode(v int) *ControlNetArgs {
	c.ResizeMode = &v
	return c
}

func (c *ControlNetArgs) SetLowvram(v bool) *ControlNetArgs {
	c.Lowvram = &v
	return c
}

func (c *ControlNetArgs) SetProcessorRes(v int) *ControlNetArgs {
	c.ProcessorRes = &v
	return c
}

func (c *ControlNetArgs) SetThresholdA(v int) *ControlNetArgs {
	c.ThresholdA = &v
	return c
}

func (c *ControlNetArgs) SetThresholdB(v int) *ControlNetArgs {
	c.ThresholdB = &v
	return c
}

func (c *ControlNetArgs) SetGuidance(v float64) *ControlNetArgs {
	c.Guidance = &v
	return c
}

func (c *ControlNetArgs) SetGuidanceStart(v float64) *ControlNetArgs {
	c.GuidanceStart = &v
	return c
}

func (c *ControlNetArgs) SetGuidanceEnd(v float64) *ControlNetArgs {
	c.GuidanceEnd = &v
	return c
}

func (c *ControlNetArgs) SetControlMode(v int) *ControlNetArgs {
	c.ControlMode = &v
	return c
}

func (c *ControlNetArgs) SetPixelPerfect(v bool) *ControlNetArgs {
	c.PixelPerfect = &v
	return c
}
