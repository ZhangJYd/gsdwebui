package gsdwebui

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestControlNet(t *testing.T) {
	c, _ := newControlNet()
	log.Infoln(c.Version())
	log.Infoln(c.ModelList())
	//dir := "D:\\code\\gsdwebui\\tmp\\"
	f1, err := os.Open("D:\\code\\gsdwebui\\tmp\\img.png")
	if err != nil {
		log.Error(err)
		return
	}
	defer f1.Close()
	//img, err := utils.ImageToBase64(f1)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}

	modules, err := c.ModuleList()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(modules)
	//wg := sync.WaitGroup{}
	//
	//for i := range modules {
	//	wg.Add(1)
	//	go func(j int) {
	//		defer wg.Done()
	//		res := &models.ControlNetDetectArgs{
	//			ControlNetModule:      modules[j],
	//			ControlNetInputImages: []string{img},
	//		}
	//		rets, err := c.Detect(res)
	//		if err != nil {
	//			log.Error(err, modules[j])
	//			return
	//		}
	//		if len(rets) == 0 {
	//			log.Error(err, modules[j])
	//			return
	//		}
	//
	//		out := path.Join(dir, modules[j]+".png")
	//		f, err := os.Create(out)
	//		if err != nil {
	//			log.Error(err)
	//			return
	//		}
	//		defer f.Close()
	//
	//		bs2, err := utils.Base64ToImage(rets[0])
	//		if err != nil {
	//			log.Error(err)
	//			return
	//		}
	//		f.Write(bs2)
	//	}(i)
	//
	//}
	//wg.Wait()

}
