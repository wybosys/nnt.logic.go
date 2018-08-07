package manager

type app struct {
	Appcfg string
	Devcfg string
}

var (
	gs_app = app{
		Appcfg: "~/app.json",
		Devcfg: "~/devops.json",
	}
)

func App() *app {
	return &gs_app
}

func (_ *app) LoadConfig() {

}

func (_ *app) Start() {

}

func (_ *app) Stop() {

}