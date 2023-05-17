package gsdwebui

import (
	"fmt"
	"gsdwebui/utils"
	"os"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStableDiffusion_Img2Img(t *testing.T) {
	dir := "D:\\code\\gsdwebui\\tmp\\"
	f1, err := os.Open("D:\\\\code\\gsdwebui\\tmp\\img.png")
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

	st, _ := NewStableDiffusion()
	payloads := NewImg2imgArgs(img.Base64).
		SetPrompt(
			"a woman with a braid smiling at the camera with a flower in the background and a blue Woven necklace on her neck, phuoc quan, a character portrait, dau-al-set,(masterpiece, sidelighting, finely detailed beautiful eyes: 1.2),best quality, masterpiece, upper body,side braids,(black eyes:1.5), Chinese Girl,Open eyes,look forward,detailed eyes,facial lighting,perfect face, <lora:gyokai:1:test>").
		SetSizeByImgInfo(img, 0.5)

	log.Info(st.ControlNetModuleList())
	log.Info(st.ControlNetModelList())
	cn := NewControlNetArgs().SetPixelPerfect(true).SetModule("openpose_face")
	canny, _ := st.FindFirstControlNetModel("openpose")
	cn.SetModel(canny).SetWeight(0.5)
	// payloads.AddControlNets(cn)

	cn = NewControlNetArgs().SetPixelPerfect(true).SetModule("hed")
	hed, _ := st.FindFirstControlNetModel("hed")
	cn.SetModel(hed).SetWeight(0.5)
	// payloads.AddControlNets(cn)
	payloads.SetDenoisingStrength(1)

	ret, err := st.Img2Img(payloads)
	if err != nil {
		log.Error(err)
		return
	}

	imgs := ret.GetImages()
	for i := range imgs {
		func(i int) {
			out := path.Join(dir, fmt.Sprintf("%v.png", i))
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
