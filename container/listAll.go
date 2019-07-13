package container

import (
	docker "github.com/helmutkemper/iotmaker.docker"
	jsonOut "github.com/helmutkemper/iotmaker.server.json"
	"net/http"
)

func NewWebContainer() webContainer {
	var ret = webContainer{}
	ret.out = jsonOut.NewJSonOut()

	return ret
}

type webContainer struct {
	out jsonOut.Out
}

func (el *webContainer) Byte() []byte {
	return el.out.Byte()
}

func (el *webContainer) ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	err, container := docker.NewContainer()
	if err != nil {
		el.out.Meta.AddError(err.Error())
		_, _ = w.Write(el.out.Byte())
		return
	}

	err, el.out.Objects = container.GetList()

	_, _ = w.Write(el.out.Byte())
}
