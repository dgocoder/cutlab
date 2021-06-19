package main

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/controllers"
	"github.com/testd/cutlab/internal/core/ports"
	authgtw "github.com/testd/cutlab/internal/gateways/auth"
	clienthandler "github.com/testd/cutlab/internal/handlers/client"
	companyhandler "github.com/testd/cutlab/internal/handlers/company"
	eventhandler "github.com/testd/cutlab/internal/handlers/event"
	locationhandler "github.com/testd/cutlab/internal/handlers/location"
	resourcehandler "github.com/testd/cutlab/internal/handlers/resource"
	servicehandler "github.com/testd/cutlab/internal/handlers/service"
	userhandler "github.com/testd/cutlab/internal/handlers/user"
	clientrepo "github.com/testd/cutlab/internal/repositories/client"
	companyrepo "github.com/testd/cutlab/internal/repositories/company"
	eventrepo "github.com/testd/cutlab/internal/repositories/event"
	locationrepo "github.com/testd/cutlab/internal/repositories/location"
	resourcerepo "github.com/testd/cutlab/internal/repositories/resource"
	servicerepo "github.com/testd/cutlab/internal/repositories/service"
	userrepo "github.com/testd/cutlab/internal/repositories/user"
	"go.uber.org/fx"
)

func AppContext() context.Context {
	return context.Background()
}
func NewFirestoreClient(ctx context.Context) *firestore.Client {
	client, _ := firestore.NewClient(ctx, "cutlab-dev")
	return client
}
func NewFirebaseApp(ctx context.Context) *firebase.App {
	app, _ := firebase.NewApp(ctx, nil)
	return app
}

func RegisterRoutes(instance *echo.Echo, companyhandler *companyhandler.HTTPHandler, clienthandler *clienthandler.HTTPHandler, eventHandler *eventhandler.HTTPHandler, locationHandler *locationhandler.HTTPHandler, resourceHandler *resourcehandler.HTTPHandler, serviceHandler *servicehandler.HTTPHandler, userHandler *userhandler.HTTPHandler) {
	instance.GET("/company/:id", companyhandler.Get)
	instance.POST("/company", companyhandler.Create)
	instance.GET("/client/:id", clienthandler.Get)
	instance.POST("/client", clienthandler.Create)
	instance.GET("/event/:id", eventHandler.Get)
	instance.POST("/event", eventHandler.Create)
	instance.GET("/location/:id", locationHandler.Get)
	instance.POST("/location", locationHandler.Create)
	instance.GET("/resource/:id", resourceHandler.Get)
	instance.POST("/resource", resourceHandler.Create)
	instance.GET("/service/:id", serviceHandler.Get)
	instance.POST("/service", serviceHandler.Create)
	instance.GET("/user/:id", userHandler.Get)
	instance.POST("/user", userHandler.Create)
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
			NewFirebaseApp,
			NewFirestoreClient,
			ConfigEcho,
			companyhandler.NewHTTPHandler,
			controllers.NewCompanyController,
			func(r *controllers.CompanyController) ports.CompanyController { return r },
			companyrepo.NewCompanyFirestoreRespository,
			func(r *companyrepo.CompanyFirestoreRespository) ports.CompanyRepository { return r },
			clienthandler.NewHTTPHandler,
			controllers.NewClientController,
			func(r *controllers.ClientController) ports.ClientController { return r },
			clientrepo.NewClientFirestoreRespository,
			func(r *clientrepo.ClientFirestoreRespository) ports.ClientRepository { return r },
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
			userhandler.NewHTTPHandler,
			controllers.NewUserController,
			func(r *controllers.UserController) ports.UserController { return r },
			userrepo.NewUserFirestoreRespository,
			func(r *userrepo.UserFirestoreRespository) ports.UserRepository { return r },
			authgtw.NewAuthFirebase,
			func(g *authgtw.AuthFirebase) ports.AuthGateway { return g },
		),
		fx.Invoke(RegisterRoutes),
	)
	app.Run()
}
