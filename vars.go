package gsdwebui

const (
	ControlNetModelList  = "/controlnet/model_list"
	ControlNetModuleList = "/controlnet/module_list"
	ControlNetDetect     = "/controlnet/detect"
	ControlNetVersion    = "/controlnet/version"
)

const (
	// task
	Img2ImgRoute            = "/sdapi/v1/img2img"
	Txt2ImgRoute            = "/sdapi/v1/txt2img"             // TODO
	ExtraSingleImageRoute   = "/sdapi/v1/extra-single-image"  // TODO
	ExtraBatchImageRoute    = "/sdapi/v1/extra-batch-images"  // TODO
	PngInfoRoute            = "/sdapi/v1/png-info"            // TODO
	InterrogateRoute        = "/sdapi/v1/interrogate"         // TODO
	InterruptRoute          = "/sdapi/v1/interrupt"           // TODO
	SkipRoute               = "/sdapi/v1/skip"                // TODO
	RefreshCheckpointsRoute = "/sdapi/v1/refresh-checkpoints" // TODO
	preprocessRoute         = "/sdapi/v1/preprocess"          // TODO

	// info
	SDModelsRoute         = "/sdapi/v1/sd-models"
	ScriptsRoute          = "/sdapi/v1/scripts"
	HypernetRoute         = "/sdapi/v1/hypernetworks"
	FaceRestoresRoute     = "/sdapi/v1/face-restorers"
	SamplersRoute         = "/sdapi/v1/samplers"
	ProgressRoute         = "/sdapi/v1/progress"
	CmdFlagsRoute         = "/sdapi/v1/cmd-flags"
	UpscalersRoute        = "/sdapi/v1/upscalers"
	RealesrganModelsRoute = "/sdapi/v1/realesrgan-models"
	PromptStylesRoute     = "/sdapi/v1/prompt-styles"
	EmbeddingsRoute       = "/sdapi/v1/embeddings"
	MemoryRoute           = "/sdapi/v1/memory"
	LoraRoute             = "/sdapi/v1/loras"

	// option
	OptionsRoute = "/sdapi/v1/options"
)
