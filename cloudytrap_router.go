package cloudytrap

package tracking

func (h *handler) initSubRoutes() {
	h._defaultHandle = CtrlIndex
	h._subRoutes = []route{
		route{pattern: "index", fn: HandleIndex},
	}
}
