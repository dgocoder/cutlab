package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/controllers"
	"github.com/testd/cutlab/internal/core/ports"
	companyhandler "github.com/testd/cutlab/internal/handlers/company"
	customerhandler "github.com/testd/cutlab/internal/handlers/customer"
	eventhandler "github.com/testd/cutlab/internal/handlers/event"
	locationhandler "github.com/testd/cutlab/internal/handlers/location"
	resourcehandler "github.com/testd/cutlab/internal/handlers/resource"
	servicehandler "github.com/testd/cutlab/internal/handlers/service"
	companyrepo "github.com/testd/cutlab/internal/repositories/company"
	customerrepo "github.com/testd/cutlab/internal/repositories/customer"
	eventrepo "github.com/testd/cutlab/internal/repositories/event"
	locationrepo "github.com/testd/cutlab/internal/repositories/location"
	resourcerepo "github.com/testd/cutlab/internal/repositories/resource"
	servicerepo "github.com/testd/cutlab/internal/repositories/service"
	"go.uber.org/fx"
)

func AppContext() context.Context {
	return context.Background()
}
func NewFirestoreClient(ctx context.Context) *firestore.Client {
	client, _ := firestore.NewClient(ctx, "cutlab-dev")
	return client
}
func RegisterRoutes(instance *echo.Echo, companyhandler *companyhandler.HTTPHandler, customerhandler *customerhandler.HTTPHandler, eventHandler *eventhandler.HTTPHandler, locationHandler *locationhandler.HTTPHandler, resourceHandler *resourcehandler.HTTPHandler, serviceHandler *servicehandler.HTTPHandler) {
	instance.GET("/company/:id", companyhandler.Get)
	instance.POST("/company", companyhandler.Create)
	instance.GET("/customer/:id", customerhandler.Get)
	instance.POST("/customer", customerhandler.Create)
	instance.GET("/event/:id", eventHandler.Get)
	instance.POST("/event", eventHandler.Create)
	instance.GET("/location/:id", locationHandler.Get)
	instance.POST("/location", locationHandler.Create)
	instance.GET("/resource/:id", resourceHandler.Get)
	instance.POST("/resource", resourceHandler.Create)
	instance.GET("/service/:id", serviceHandler.Get)
	instance.POST("/service", serviceHandler.Create)
}
func ConfigEcho(lc fx.Lifecycle) *echo.Echo {
	echo := echo.New()
	lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go echo.Start(":8080")
			return nil
		},
		OnStop: func(context context.Context) error {
			return echo.Close()
		},
	})
	return echo
}
func main() {
	app := fx.New(
		fx.Provide(
			AppContext,
			NewFirestoreClient,
			ConfigEcho,
			companyhandler.NewHTTPHandler,
			controllers.NewCompanyController,
			func(r *controllers.CompanyController) ports.CompanyController { return r },
			companyrepo.NewCompanyFirestoreRespository,
			func(r *companyrepo.CompanyFirestoreRespository) ports.CompanyRepository { return r },
			customerhandler.NewHTTPHandler,
			controllers.NewCustomerController,
			func(r *controllers.CustomerController) ports.CustomerController { return r },
			customerrepo.NewCustomerFirestoreRespository,
			func(r *customerrepo.CustomerFirestoreRespository) ports.CustomerRepository { return r },
			eventhandler.NewHTTPHandler,
			controllers.NewEventController,
			func(r *controllers.EventController) ports.EventController { return r },
			eventrepo.NewEventFirestoreRespository,
			func(r *eventrepo.EventFirestoreRespository) ports.EventRepository { return r },
			locationhandler.NewHTTPHandler,
			controllers.NewLocationController,
			func(r *controllers.LocationController) ports.LocationController { return r },
			locationrepo.NewLocationFirestoreRespository,
			func(r *locationrepo.LocationFirestoreRespository) ports.LocationRepository { return r },
			resourcehandler.NewHTTPHandler,
			controllers.NewResourceController,
			func(r *controllers.ResourceController) ports.ResourceController { return r },
			resourcerepo.NewResourceFirestoreRespository,
			func(r *resourcerepo.ResourceFirestoreRespository) ports.ResourceRepository { return r },
			servicehandler.NewHTTPHandler,
			controllers.NewServiceController,
			func(r *controllers.ServiceController) ports.ServiceController { return r },
			servicerepo.NewServiceFirestoreRespository,
			func(r *servicerepo.ServiceFirestoreRespository) ports.ServiceRepository { return r },
		),
		fx.Invoke(RegisterRoutes),
	)
	app.Run()
}
