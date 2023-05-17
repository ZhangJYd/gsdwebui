package models

import (
	"encoding/json"
	"gsdwebui/utils"
)

type SDTxt2imgArgs struct {
	EnableHr          *bool    `json:"enable_hr,omitempty"`
	DenoisingStrength *float64 `json:"denoising_strength,omitempty"`
	FirstphaseWidth   *int     `json:"firstphase_width,omitempty"`
	FirstphaseHeight  *int     `json:"firstphase_height,omitempty"`
	HrScale           *float64 `json:"hr_scale,omitempty"`
	HrUpscaler        string   `json:"hr_upscaler,omitempty"`
	HrSecondPassSteps *int     `json:"hr_second_pass_steps,omitempty"`
	HrResizeX         *int     `json:"hr_resize_x,omitempty"`
	HrResizeY         *int     `json:"hr_resize_y,omitempty"`
	SMinUncond        *int     `json:"s_min_uncond,omitempty"`

	SDCommonArgs
}

func (s *SDTxt2imgArgs) SetEnableHr(v bool) *SDTxt2imgArgs {
	s.EnableHr = &v
	return s
}

func (s *SDTxt2imgArgs) SetDenoisingStrength(v int) *SDTxt2imgArgs {
	s.FirstphaseWidth = &v
	return s
}

func (s *SDTxt2imgArgs) SetFirstphaseHeight(v int) *SDTxt2imgArgs {
	s.FirstphaseHeight = &v
	return s
}

func (s *SDTxt2imgArgs) SetHrUpscaler(v string) *SDTxt2imgArgs {
	s.HrUpscaler = v
	return s
}

func (s *SDTxt2imgArgs) SetHrSecondPassSteps(v int) *SDTxt2imgArgs {
	s.HrSecondPassSteps = &v
	return s
}

func (s *SDTxt2imgArgs) SetHrResizeX(v int) *SDTxt2imgArgs {
	s.HrResizeX = &v
	return s
}

func (s *SDTxt2imgArgs) SetHrResizeY(v int) *SDTxt2imgArgs {
	s.HrResizeY = &v
	return s
}

func (s *SDTxt2imgArgs) SetPrompt(v string) *SDTxt2imgArgs {
	s.Prompt = v
	return s
}

func (s *SDTxt2imgArgs) AddStyles(v ...string) *SDTxt2imgArgs {
	s.Styles = append(s.Styles, v...)
	return s
}

func (s *SDTxt2imgArgs) SetSeed(v int) *SDTxt2imgArgs {
	s.Seed = &v
	return s
}

func (s *SDTxt2imgArgs) SetSubseed(v int) *SDTxt2imgArgs {
	s.Subseed = &v
	return s
}

func (s *SDTxt2imgArgs) SetSubseedStrength(v int) *SDTxt2imgArgs {
	s.SubseedStrength = &v
	return s
}

func (s *SDTxt2imgArgs) SetSeedResizeFromH(v int) *SDTxt2imgArgs {
	s.SeedResizeFromH = &v
	return s
}

func (s *SDTxt2imgArgs) SetSeedResizeFromW(v int) *SDTxt2imgArgs {
	s.SeedResizeFromW = &v
	return s
}

func (s *SDTxt2imgArgs) SetSamplerName(v string) *SDTxt2imgArgs {
	s.SamplerName = v
	return s
}

func (s *SDTxt2imgArgs) SetSamplerIndex(v string) *SDTxt2imgArgs {
	s.SamplerIndex = v
	return s
}

func (s *SDTxt2imgArgs) SetBatchSize(v int) *SDTxt2imgArgs {
	s.BatchSize = &v
	return s
}

func (s *SDTxt2imgArgs) SetNIter(v int) *SDTxt2imgArgs {
	s.NIter = &v
	return s
}

func (s *SDTxt2imgArgs) SetSteps(v int) *SDTxt2imgArgs {
	s.Steps = &v
	return s
}

func (s *SDTxt2imgArgs) SetCfgScale(v float64) *SDTxt2imgArgs {
	s.CfgScale = &v
	return s
}

func (s *SDTxt2imgArgs) SetSize(width int, high int) *SDTxt2imgArgs {
	s.Height = &high
	s.Width = &width
	return s
}

func (s *SDTxt2imgArgs) SetSizeByImgInfo(i *utils.ImageInfo, rate float64) *SDTxt2imgArgs {
	s.Height = utils.P(int(float64(i.Height) * rate))
	s.Width = utils.P(int(float64(i.Width) * rate))
	return s
}

func (s *SDTxt2imgArgs) SetRestoreFaces(v bool) *SDTxt2imgArgs {
	s.RestoreFaces = &v
	return s
}

func (s *SDTxt2imgArgs) SetTiling(v bool) *SDTxt2imgArgs {
	s.Tiling = &v
	return s
}

func (s *SDTxt2imgArgs) SetDoNotSaveSamples(v bool) *SDTxt2imgArgs {
	s.DoNotSaveSamples = &v
	return s
}

func (s *SDTxt2imgArgs) SetDoNotSaveGrid(v bool) *SDTxt2imgArgs {
	s.DoNotSaveGrid = &v
	return s
}

func (s *SDTxt2imgArgs) SetNegativePrompt(v string) *SDTxt2imgArgs {
	s.NegativePrompt = v
	return s
}

func (s *SDTxt2imgArgs) SetEta(v float64) *SDTxt2imgArgs {
	s.Eta = &v
	return s
}

func (s *SDTxt2imgArgs) SetSMinUncond(v int) *SDTxt2imgArgs {
	s.SMinUncond = &v
	return s
}

func (s *SDTxt2imgArgs) SetSChurn(v int) *SDTxt2imgArgs {
	s.SChurn = &v
	return s
}

func (s *SDTxt2imgArgs) SetSTmax(v int) *SDTxt2imgArgs {
	s.STmax = &v
	return s
}

func (s *SDTxt2imgArgs) SetSTmin(v int) *SDTxt2imgArgs {
	s.STmin = &v
	return s
}

func (s *SDTxt2imgArgs) SetSNoise(v int) *SDTxt2imgArgs {
	s.SNoise = &v
	return s
}

func (s *SDTxt2imgArgs) SetOverrideSettings(k string, v any) *SDTxt2imgArgs {
	if s.OverrideSettings == nil {
		s.OverrideSettings = make(map[string]any)
	}
	s.OverrideSettings[k] = v
	return s
}

func (s *SDTxt2imgArgs) SetOverrideSettingsRestoreAfterwards(v bool) *SDTxt2imgArgs {
	s.OverrideSettingsRestoreAfterwards = &v
	return s
}

func (s *SDTxt2imgArgs) SetScriptArgs(k string, v any) *SDTxt2imgArgs {
	if s.ScriptArgs == nil {
		s.ScriptArgs = make(map[string]any)
	}
	s.ScriptArgs[k] = v
	return s
}

func (s *SDTxt2imgArgs) SetSendImages(v bool) *SDTxt2imgArgs {
	s.SendImages = &v
	return s
}

func (s *SDTxt2imgArgs) SetSaveImages(v bool) *SDTxt2imgArgs {
	s.SaveImages = &v
	return s
}

func (s *SDTxt2imgArgs) SetScriptName(v string) *SDTxt2imgArgs {
	s.ScriptName = v
	return s
}

func (s *SDTxt2imgArgs) SetAlwaysonScripts(k string, v any) *SDTxt2imgArgs {
	if s.AlwaysonScripts == nil {
		s.AlwaysonScripts = make(map[string]any)
	}
	s.AlwaysonScripts[k] = v
	return s
}

func (s *SDTxt2imgArgs) AddControlNets(cs ...*ControlNetArgs) *SDTxt2imgArgs {
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

func (s *SDTxt2imgArgs) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}
