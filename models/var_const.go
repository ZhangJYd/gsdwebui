package models

const (
	UpscalerNone                  = "None"
	UpscalerLanczos               = "Lanczos"
	UpscalerNearest               = "Nearest"
	UpscalerLDSR                  = "LDSR"
	UpscalerBSRGAN                = "BSRGAN"
	UpscalerESRGAN_4x             = "ESRGAN_4x"
	UpscalerR_ESRGAN_General_4xV3 = "R-ESRGAN General 4xV3"
	UpscalerScuNET_GAN            = "ScuNET GAN"
	UpscalerScuNET_PSNR           = "ScuNET PSNR"
	UpscalerSwinIR_4x             = "SwinIR 4x"
)

const (
	HiResUpscalerNone                         = "None"
	HiResUpscalerNoneLatent                   = "Latent"
	HiResUpscalerNoneLatentAntialiased        = "Latent (antialiased)"
	HiResUpscalerNoneLatentBicubic            = "Latent (bicubic)"
	HiResUpscalerNoneLatentBicubicAntialiased = "Latent (bicubic antialiased)"
	HiResUpscalerNoneLatentNearest            = "Latent (nearist)"
	HiResUpscalerNoneLatentNearestExact       = "Latent (nearist-exact)"
	HiResUpscalerNoneLanczos                  = "Lanczos"
	HiResUpscalerNoneNearest                  = "Nearest"
	HiResUpscalerNoneESRGAN_4x                = "ESRGAN_4x"
	HiResUpscalerNoneLDSR                     = "LDSR"
	HiResUpscalerNoneScuNET_GAN               = "ScuNET GAN"
	HiResUpscalerNoneScuNET_PSNR              = "ScuNET PSNR"
	HiResUpscalerNoneSwinIR_4x                = "SwinIR 4x"
)
