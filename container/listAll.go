package container

import (
	docker "github.com/helmutkemper/iotmaker.docker"
	jsonOut "github.com/helmutkemper/iotmaker.server.json"
	"net/http"
)

func NewWebContainer() WebContainer {
	var ret = WebContainer{}
	ret.out = jsonOut.NewJSonOut()

	return ret
}

type WebContainer struct {
	out jsonOut.Out
}

func (el *WebContainer) Byte() []byte {
	return el.out.Byte()
}

func (el *WebContainer) ListAll(w http.ResponseWriter, r *http.Request) {

	err, container := docker.NewContainerStt()
	if err != nil {
		el.out.Meta.AddError(err.Error())
		_, _ = w.Write(el.out.Byte())
		return
	}

	err, el.out.Objects = container.GetList()

	_, _ = w.Write(el.out.Byte())
}
