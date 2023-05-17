package main

import (
	"fmt"
	"gsdwebui"
	"gsdwebui/utils"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

func main() {
	dir := "tmp"
	f1, err := os.Open("tmp/1.jpg")
	if err != nil {
		log.Error(err)
		return
	}
	defer f1.Close()
	img, err := utils.GetImageInfo(f1)
	if err != nil {
		log.Error(err)
		return
	}

	st, err := gsdwebui.NewStableDiffusion()
	if err != nil {
		log.Error(err)
		return
	}

	depthControl := gsdwebui.NewControlNetArgs()
	depthModel, _ := st.FindFirstControlNetModel("anime")
	depthControl.SetModel(depthModel).
		SetModule("lineart_anime_denoise").
		SetPixelPerfect(true).
		SetWeight(0.5)

	ret, err := st.Img2Img(
		gsdwebui.NewImg2imgArgs(img.Base64).
			SetPrompt("a woman with a braid smiling at the camera with a flower in the background and a blue Woven necklace on her neck, phuoc quan, a character portrait, dau-al-set,(masterpiece, sidelighting, finely detailed beautiful eyes: 1.2),best quality, masterpiece, upper body,side braids,(black eyes:1.5), Chinese Girl,Open eyes,look forward,detailed eyes,facial lighting,perfect face, <lora:twitch_emotes_v1.3:1:test>,earring,White clothes,Reveal the neck").
			SetNegativePrompt("EasyNegative, badhandv4,(worst quality, low quality,collar:1.4)").
			SetSamplerName("DPM++ 2M Karras").
			SetSize(448, 597).
			SetCfgScale(9).
			SetDenoisingStrength(1).
			SetSeed(3701119662).
			SetSteps(40).
			AddControlNets(depthControl),
	)
	if err != nil {
		log.Error(err)
		return
	}

	imgs := ret.GetImages()
	for i := range imgs {
		func(i int) {
			out := path.Join(dir, fmt.Sprintf("img2img%v.png", i))
			f, err := os.Create(out)
			if err != nil {
				log.Error(err)
				return
			}
			defer f.Close()
			bs2, err := utils.Base64ToImage(imgs[0])
			if err != nil {
				log.Error(err)
				return
			}
			f.Write(bs2)
		}(i)
	}

	depthControl.SetInputImage(img.Base64)

	ret, err = st.Txt2Img(
		gsdwebui.NewTxt2imgArgs().
			SetPrompt("a woman with a braid smiling at the camera with a flower in the background and a blue Woven necklace on her neck, phuoc quan, a character portrait, dau-al-set,(masterpiece, sidelighting, finely detailed beautiful eyes: 1.2),best quality, masterpiece, upper body,side braids,(black eyes:1.5), Chinese Girl,Open eyes,look forward,detailed eyes,facial lighting,perfect face, <lora:twitch_emotes_v1.3:1:test>,earring,White clothes,Reveal the neck").
			SetNegativePrompt("EasyNegative, badhandv4,(worst quality, low quality,collar:1.4)").
			SetSamplerName("DPM++ 2M Karras").
			SetSize(448, 597).
			SetCfgScale(9).
			SetSeed(3701119662).
			SetSteps(40),
		//AddControlNets(depthControl),
	)
	if err != nil {
		log.Error(err)
		return
	}
	imgs = ret.GetImages()
	for i := range imgs {
		func(i int) {
			out := path.Join(dir, fmt.Sprintf("txt2img%v.png", i))
			f, err := os.Create(out)
			if err != nil {
				log.Error(err)
				return
			}
			defer f.Close()
			bs2, err := utils.Base64ToImage(imgs[0])
			if err != nil {
				log.Error(err)
				return
			}
			f.Write(bs2)
		}(i)
	}
}
