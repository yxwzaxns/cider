package server

func Init() {
	// config := config.GetConfig()
	r := NewRouter()
	r.Run()
}
