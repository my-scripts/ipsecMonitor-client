package rpc

type Handler struct {
}

func (this *Handler) RestartIpsec(args *Args, reply *Status) error {
	reply.Succ = true
	// var us db.UploadSetting
	// if this.DB.First(&us).Error != nil {
	// 	cfg := config.Config{}
	// 	err := cfg.Load(filepath.Join("conf", "tramagent.json"))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*reply = TusReq{Addr: cfg.Tus.Addr, Port: cfg.Tus.Port, UploadTime: -1, Capacity: 30}
	// 	return nil
	//
	// }
	//
	// *reply = TusReq{Addr: us.Addr, Port: us.Port, UploadTime: us.UploadTime, Capacity: us.Capacity}
	return nil
}
