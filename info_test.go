package gsdwebui

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStableDiffusion_GetOptions(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetOptions()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetOptions2(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetOptions2()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetSDModels(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetSDModels()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetScripts(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetScripts()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetHypenetworks(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetHypenetworks()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetFaceRestores(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetFaceRestores()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetSamplers(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetSamplers()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetProgress(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetProgress()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetCmdFlags(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetCmdFlags()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetUpscalers(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetUpscalers()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetRealesrganModels(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetRealesrganModels()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetPromptStyles(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetPromptStyles()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetEmbeddings(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetEmbeddings()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetMemoryInfo(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetMemoryInfo()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}

func TestStableDiffusion_GetLoras(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetLoras()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%+v", ret)
}
