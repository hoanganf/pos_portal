package src

import (
	"errors"
	"github.com/gin-contrib/multitemplate"
	"github.com/hoanganf/pos_portal/src/application"
	"github.com/hoanganf/pos_portal/src/infrastructure/client"
	"github.com/hoanganf/pos_portal/src/infrastructure/persistence"
	"os"
	"strconv"
)

const (
	PageCashier    = "cashier"
	PageRestaurant = "restaurant"
	CommonTitle    = "title"
	CommonHeader   = "header"
	CommonFooter   = "footer"
	TemplatePath   = "./templates"
)

type Bean struct {
	PageLayouts       map[string]string
	CommonLayouts     map[string]string
	CashierService    *application.CashierService
	RestaurantService *application.RestaurantService
	PortalApiService  *application.PortalApiService
}

func InitBean() (*Bean, error) {
	userTimeout, err := strconv.Atoi(getEnvWithDefault("POS_USER_TIMEOUT", "5000"))
	if err != nil {
		return nil, err
	}
	userClient := client.NewUserClient(
		getEnvWithDefault("POS_USER_HOST", "http://localhost:8080"),
		userTimeout,
	)

	serverApiTimeout, err := strconv.Atoi(getEnvWithDefault("POS_SERVER_API_TIMEOUT", "5000"))
	if err != nil {
		return nil, err
	}

	serverApiClient := client.NewServerApiClient(
		getEnvWithDefault("POS_SERVER_API_HOST", "http://localhost:8080"),
		serverApiTimeout,
	)

	localApiTimeout, err := strconv.Atoi(getEnvWithDefault("POS_LOCAL_API_TIMEOUT", "5000"))
	if err != nil {
		return nil, err
	}
	localApiClient := client.NewLocalApiClient(
		getEnvWithDefault("POS_LOCAL_API_HOST", "http://localhost:8081"),
		localApiTimeout,
	)

	loginRepository := persistence.NewLoginRepository(userClient)
	serverApiRepository := persistence.NewServerApiRepository(serverApiClient)
	localApiRepository := persistence.NewLocalApiRepository(localApiClient)

	commonService := application.NewCommonService(
		getEnvWithDefault("POS_LOGIN_URL", "http://localhost:8080/login"),
		getEnvWithDefault("POS_DESIGN_URL", "http://localhost/pos/pos-lib"),
		getEnvWithDefault("POS_IMAGE_URL", "http://localhost"),
		getEnvWithDefault("POS_LOGIN_TOKEN", "pos_access_token"),
		loginRepository)

	cashierService := application.NewCashierService(commonService, serverApiRepository)
	restaurantService := application.NewRestaurantService(commonService, serverApiRepository)
	portalApiService := application.NewPortalApiService(commonService, localApiRepository)

	// html layout
	pageLayouts := make(map[string]string)
	pageLayouts[PageCashier] = TemplatePath + "/page_cashier.html"
	pageLayouts[PageRestaurant] = TemplatePath + "/page_restaurant.html"

	//common
	commonLayouts := make(map[string]string)
	commonLayouts[CommonHeader] = TemplatePath + "/commons/header.html"
	commonLayouts[CommonFooter] = TemplatePath + "/commons/footer.html"
	commonLayouts[CommonTitle] = TemplatePath + "/commons/title.html"

	return &Bean{
		PageLayouts:       pageLayouts,
		CommonLayouts:     commonLayouts,
		CashierService:    cashierService,
		RestaurantService: restaurantService,
		PortalApiService:  portalApiService}, nil
}

func (bean *Bean) LoadTemplates() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//cashierPage
	render.AddFromFiles(PageCashier,
		bean.PageLayouts[PageCashier],
		bean.CommonLayouts[CommonTitle],
		bean.CommonLayouts[CommonHeader],
		bean.CommonLayouts[CommonFooter])

	//restaurantPage
	render.AddFromFiles(PageRestaurant, bean.PageLayouts[PageRestaurant])
	return render
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}

func getEnvRequired(name string) (string, error) {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env, nil
	}
	return "", errors.New("not found env: " + name)
}
