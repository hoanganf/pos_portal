package resource

type PageResource struct {
	ErrorMessage string
	Message      string
	IsCashier    bool
	IsOrder      bool
	Areas        []interface{}
	Restaurants  []interface{}
	LoginUrl     string
	DesignUrl    string
	ImageUrl     string
	PageTitle    string
	Origin       string
}
