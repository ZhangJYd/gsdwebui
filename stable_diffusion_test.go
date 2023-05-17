package gsdwebui

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStableDiffusion_SetOptions(t *testing.T) {
	st, _ := NewStableDiffusion()
	err := st.SetOptionsAny(map[string]any{
		"sd_checkpoint_cache": 1,
	})
	if err != nil {
		log.Error(err)
		return
	}
}

func TestStableDiffusion_SetSDModelCheckpoint(t *testing.T) {
	st, _ := NewStableDiffusion()
	ret, err := st.GetSDModels()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(ret)
	err = st.SetSDModelCheckpoint(&ret[0])
	log.Info(err)
}
