package models

type SDCommonArgs struct {
	Prompt                            string         `json:"prompt,omitempty"`
	NegativePrompt                    string         `json:"negative_prompt,omitempty"`
	Styles                            []string       `json:"styles,omitempty"`
	Seed                              *int           `json:"seed,omitempty"`
	Subseed                           *int           `json:"subseed,omitempty"`
	SubseedStrength                   *int           `json:"subseed_strength,omitempty"`
	SeedResizeFromH                   *int           `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   *int           `json:"seed_resize_from_w,omitempty"`
	SamplerIndex                      string         `json:"sampler_index,omitempty"`
	SamplerName                       string         `json:"sampler_name,omitempty"`
	BatchSize                         *int           `json:"batch_size,omitempty"`
	NIter                             *int           `json:"n_iter,omitempty"`
	Steps                             *int           `json:"steps,omitempty"`
	CfgScale                          *float64       `json:"cfg_scale,omitempty"`
	Width                             *int           `json:"width,omitempty"`
	Height                            *int           `json:"height,omitempty"`
	RestoreFaces                      *bool          `json:"restore_faces,omitempty"`
	Tiling                            *bool          `json:"tiling,omitempty"`
	DoNotSaveSamples                  *bool          `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid                     *bool          `json:"do_not_save_grid,omitempty"`
	Eta                               *float64       `json:"eta,omitempty"`
	SChurn                            *int           `json:"s_churn,omitempty"`
	STmax                             *int           `json:"s_tmax,omitempty"`
	STmin                             *int           `json:"s_tmin,omitempty"`
	SNoise                            *int           `json:"s_noise,omitempty"`
	OverrideSettings                  map[string]any `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards *bool          `json:"override_settings_restore_afterwards,omitempty"`
	ScriptName                        string         `json:"script_name,omitempty"`
	ScriptArgs                        map[string]any `json:"script_args,omitempty"`
	SendImages                        *bool          `json:"send_images,omitempty"`
	SaveImages                        *bool          `json:"save_images,omitempty"`
	AlwaysonScripts                   map[string]any `json:"alwayson_scripts,omitempty"`
}
