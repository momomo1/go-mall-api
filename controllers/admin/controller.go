package admin

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func whereAnd(where *string) {
	if *where != "" {
		*where += "and "
	}
}
