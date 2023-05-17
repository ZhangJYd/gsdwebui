package models

import (
	"encoding/json"
)

type SDModelInfo struct {
	Title     string `json:"title"`
	ModelName string `json:"modelName"`
	Hash      string `json:"hash"`
	Sha256    string `json:"sha256"`
	Filename  string `json:"filename"`
	Config    string `json:"config"`
}

type ScriptsInfo struct {
	Txt2Img []string `json:"txt2img"`
	Img2Img []string `json:"img2img"`
}

type HypernetworksInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type FaceRestoresInfo struct {
	Name   string `json:"name"`
	CmdDir string `json:"cmd_dir"`
}

type SamplerInfo struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`
	Options struct {
		Scheduler              string `json:"scheduler"`
		DiscardNextToLastSigma string `json:"discard_next_to_last_sigma"`
	} `json:"options"`
}

type LoraInfo struct {
	Name     string         `json:"name"`
	Alias    string         `json:"alias"`
	Path     string         `json:"path"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

func (s *LoraInfo) String() string {
	tmp := *s
	tmp.Metadata = nil
	data, _ := json.Marshal(tmp)
	return string(data)
}

type ProgressInfo struct {
	Progress    float64 `json:"progress"`
	EtaRelative float64 `json:"eta_relative"`
	State       struct {
		Skipped       bool   `json:"skipped"`
		Interrupted   bool   `json:"interrupted"`
		Job           string `json:"job"`
		JobCount      int    `json:"job_count"`
		JobTimestamp  string `json:"job_timestamp"`
		JobNo         int    `json:"job_no"`
		SamplingStep  int    `json:"sampling_step"`
		SamplingSteps int    `json:"sampling_steps"`
	} `json:"state"`
	CurrentImage string `json:"current_image"`
	Textinfo     string `json:"textinfo"`
}

type CmdFlagsInfo struct {
	F                             bool   `json:"f"`
	UpdateAllExtensions           bool   `json:"update_all_extensions"`
	SkipPythonVersionCheck        bool   `json:"skip_python_version_check"`
	SkipTorchCudaTest             bool   `json:"skip_torch_cuda_test"`
	ReinstallXformers             bool   `json:"reinstall_xformers"`
	ReinstallTorch                bool   `json:"reinstall_torch"`
	UpdateCheck                   bool   `json:"update_check"`
	Tests                         any    `json:"tests"`
	NoTests                       bool   `json:"no_tests"`
	SkipInstall                   bool   `json:"skip_install"`
	DataDir                       string `json:"data_dir"`
	Config                        string `json:"config"`
	Ckpt                          string `json:"ckpt"`
	CkptDir                       any    `json:"ckpt_dir"`
	VaeDir                        any    `json:"vae_dir"`
	GfpganDir                     string `json:"gfpgan_dir"`
	GfpganModel                   any    `json:"gfpgan_model"`
	NoHalf                        bool   `json:"no_half"`
	NoHalfVae                     bool   `json:"no_half_vae"`
	NoProgressbarHiding           bool   `json:"no_progressbar_hiding"`
	MaxBatchCount                 int    `json:"max_batch_count"`
	EmbeddingsDir                 string `json:"embeddings_dir"`
	TextualInversionTemplatesDir  string `json:"textual_inversion_templates_dir"`
	HypernetworkDir               string `json:"hypernetwork_dir"`
	LocalizationsDir              string `json:"localizations_dir"`
	AllowCode                     bool   `json:"allow_code"`
	Medvram                       bool   `json:"medvram"`
	Lowvram                       bool   `json:"lowvram"`
	Lowram                        bool   `json:"lowram"`
	AlwaysBatchCondUncond         bool   `json:"always_batch_cond_uncond"`
	UnloadGfpgan                  bool   `json:"unload_gfpgan"`
	Precision                     string `json:"precision"`
	UpcastSampling                bool   `json:"upcast_sampling"`
	Share                         bool   `json:"share"`
	Ngrok                         any    `json:"ngrok"`
	NgrokRegion                   string `json:"ngrok_region"`
	EnableInsecureExtensionAccess bool   `json:"enable_insecure_extension_access"`
	CodeformerModelsPath          string `json:"codeformer_models_path"`
	GfpganModelsPath              string `json:"gfpgan_models_path"`
	EsrganModelsPath              string `json:"esrgan_models_path"`
	BsrganModelsPath              string `json:"bsrgan_models_path"`
	RealesrganModelsPath          string `json:"realesrgan_models_path"`
	ClipModelsPath                any    `json:"clip_models_path"`
	Xformers                      bool   `json:"xformers"`
	ForceEnableXformers           bool   `json:"force_enable_xformers"`
	XformersFlashAttention        bool   `json:"xformers_flash_attention"`
	Deepdanbooru                  bool   `json:"deepdanbooru"`
	OptSplitAttention             bool   `json:"opt_split_attention"`
	OptSubQuadAttention           bool   `json:"opt_sub_quad_attention"`
	SubQuadQChunkSize             int    `json:"sub_quad_q_chunk_size"`
	SubQuadKvChunkSize            any    `json:"sub_quad_kv_chunk_size"`
	SubQuadChunkThreshold         any    `json:"sub_quad_chunk_threshold"`
	OptSplitAttentionInvokeai     bool   `json:"opt_split_attention_invokeai"`
	OptSplitAttentionV1           bool   `json:"opt_split_attention_v1"`
	OptSdpAttention               bool   `json:"opt_sdp_attention"`
	OptSdpNoMemAttention          bool   `json:"opt_sdp_no_mem_attention"`
	DisableOptSplitAttention      bool   `json:"disable_opt_split_attention"`
	DisableNanCheck               bool   `json:"disable_nan_check"`
	UseCPU                        []any  `json:"use_cpu"`
	Listen                        bool   `json:"listen"`
	Port                          any    `json:"port"`
	ShowNegativePrompt            bool   `json:"show_negative_prompt"`
	UIConfigFile                  string `json:"ui_config_file"`
	HideUIDirConfig               bool   `json:"hide_ui_dir_config"`
	FreezeSettings                bool   `json:"freeze_settings"`
	UISettingsFile                string `json:"ui_settings_file"`
	GradioDebug                   bool   `json:"gradio_debug"`
	GradioAuth                    any    `json:"gradio_auth"`
	GradioAuthPath                any    `json:"gradio_auth_path"`
	GradioImg2ImgTool             any    `json:"gradio_img2img_tool"`
	GradioInpaintTool             any    `json:"gradio_inpaint_tool"`
	OptChannelslast               bool   `json:"opt_channelslast"`
	StylesFile                    string `json:"styles_file"`
	Autolaunch                    bool   `json:"autolaunch"`
	Theme                         any    `json:"theme"`
	UseTextboxSeed                bool   `json:"use_textbox_seed"`
	DisableConsoleProgressbars    bool   `json:"disable_console_progressbars"`
	EnableConsolePrompts          bool   `json:"enable_console_prompts"`
	VaePath                       any    `json:"vae_path"`
	DisableSafeUnpickle           bool   `json:"disable_safe_unpickle"`
	API                           bool   `json:"api"`
	APIAuth                       any    `json:"api_auth"`
	APILog                        bool   `json:"api_log"`
	Nowebui                       bool   `json:"nowebui"`
	UIDebugMode                   bool   `json:"ui_debug_mode"`
	DeviceID                      any    `json:"device_id"`
	Administrator                 bool   `json:"administrator"`
	CorsAllowOrigins              any    `json:"cors_allow_origins"`
	CorsAllowOriginsRegex         any    `json:"cors_allow_origins_regex"`
	TLSKeyfile                    any    `json:"tls_keyfile"`
	TLSCertfile                   any    `json:"tls_certfile"`
	DisableTLSVerify              any    `json:"disable_tls_verify"`
	ServerName                    any    `json:"server_name"`
	GradioQueue                   bool   `json:"gradio_queue"`
	NoGradioQueue                 bool   `json:"no_gradio_queue"`
	SkipVersionCheck              bool   `json:"skip_version_check"`
	NoHashing                     bool   `json:"no_hashing"`
	NoDownloadSdModel             bool   `json:"no_download_sd_model"`
	AddnetMaxModelCount           int    `json:"addnet_max_model_count"`
	ControlnetDir                 any    `json:"controlnet_dir"`
	ControlnetAnnotatorModelsPath any    `json:"controlnet_annotator_models_path"`
	NoHalfControlnet              any    `json:"no_half_controlnet"`
	LdsrModelsPath                string `json:"ldsr_models_path"`
	LoraDir                       string `json:"lora_dir"`
	ScunetModelsPath              string `json:"scunet_models_path"`
	SwinirModelsPath              string `json:"swinir_models_path"`
}

type UpscalerInfo struct {
	Name      string  `json:"name"`
	ModelName string  `json:"model_name"`
	ModelPath string  `json:"model_path"`
	ModelURL  string  `json:"model_url"`
	Scale     float64 `json:"scale"`
}

type RealesrganModelInfo struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Scale int    `json:"scale"`
}

type PromptStyleInfo struct {
	Name           string `json:"name"`
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
}

type EmbeddingInfo struct {
	Step             int    `json:"step"`
	SdCheckpoint     string `json:"sd_checkpoint"`
	SdCheckpointName string `json:"sd_checkpoint_name"`
	Shape            int    `json:"shape"`
	Vectors          int    `json:"vectors"`
}

type EmbeddingsInfo struct {
	Loaded  map[string]EmbeddingInfo `json:"loaded"`
	Skipped map[string]EmbeddingInfo `json:"skipped"`
}

type MemoryInfo struct {
	RAM struct {
		Free  float64 `json:"free"`
		Used  int     `json:"used"`
		Total float64 `json:"total"`
	} `json:"ram"`
	Cuda struct {
		System struct {
			Free  int64 `json:"free"`
			Used  int64 `json:"used"`
			Total int64 `json:"total"`
		} `json:"system"`
		Active struct {
			Current int64 `json:"current"`
			Peak    int64 `json:"peak"`
		} `json:"active"`
		Allocated struct {
			Current int64 `json:"current"`
			Peak    int64 `json:"peak"`
		} `json:"allocated"`
		Reserved struct {
			Current int64 `json:"current"`
			Peak    int64 `json:"peak"`
		} `json:"reserved"`
		Inactive struct {
			Current int `json:"current"`
			Peak    int `json:"peak"`
		} `json:"inactive"`
		Events struct {
			Retries int `json:"retries"`
			Oom     int `json:"oom"`
		} `json:"events"`
	} `json:"cuda"`
}
