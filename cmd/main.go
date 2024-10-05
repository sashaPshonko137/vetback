package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"vetback/internal/api/animal"
	"vetback/internal/api/appointment"
	"vetback/internal/api/diagnosis"
	"vetback/internal/api/treatment"
	"vetback/internal/api/user"
	"vetback/internal/config"
	"vetback/internal/repo/storage"
	animalService "vetback/internal/service/animal"
	appointmentService "vetback/internal/service/appointment"
	diadnosisService "vetback/internal/service/diagnosis"
	treatmentService "vetback/internal/service/treatment"
	userService "vetback/internal/service/user"
	"vetback/metrics"
)

// @title           VET API
// @version         14.89
// @description     This is a server for owners and vets

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @in   header
// @name Authorization

func main() {
	// init logger
	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// init config
	cfg := config.NewConfig()

	// init storage
	bd, err := storage.NewStorage(cfg.DBUrl)
	if err != nil {
		panic(err)
	}

	// init service
	userServ := userService.NewUserService(bd, logger, *cfg)
	diagnosisServ := diadnosisService.NewDiagnosisService(bd, logger, userServ)
	animalServ := animalService.NewAnimalService(bd, logger, userServ, diagnosisServ)
	treatmentServ := treatmentService.NewTreatmentService(bd, logger, userServ, animalServ, diagnosisServ)
	appointmentServ := appointmentService.NewAppointmentService(bd, logger, userServ, animalServ, diagnosisServ)

	// init api
	userApi := user.NewUserApi(userServ, logger)
	animalApi := animal.NewAnimalApi(animalServ, logger, userApi)
	diagnosisApi := diagnosis.NewDiagnosisApi(diagnosisServ, logger, userApi)
	treatmentApi := treatment.NewTreatmentApi(treatmentServ, logger, userApi)
	appointmentApi := appointment.NewAppointmentApi(appointmentServ, logger, userApi)

	// init router
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(metrics.MetricsMiddleware)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), // URL для получения спецификации
	))

	router.Post("/animal", animalApi.NewCreate())
	router.Get("/animal/{id}", animalApi.NewGet())
	router.Get("/animal", animalApi.NewGetMany())
	router.Put("/animal/{id}", animalApi.NewUpdate())
	router.Delete("/animal/{id}", animalApi.NewDelete())
	router.Get("/story/{id}", animalApi.SendPDF(appointmentServ, treatmentServ))

	router.Post("/sign-up", userApi.NewSignUp())
	router.Post("/sign-in", userApi.NewSignIn())
	router.Get("/sign-out", userApi.NewSignOut())
	router.Get("/session-info", userApi.NewGetSessionInfo())

	router.Get("/user/{id}", userApi.NewGet())
	router.Get("/user", userApi.NewGetMany())
	router.Put("/user/{id}", animalApi.NewUpdate())
	router.Delete("/user", userApi.NewDelete())

	router.Post("/diagnosis", diagnosisApi.NewCreate())
	router.Get("/diagnosis/{id}", diagnosisApi.NewGet())
	router.Get("/diagnosis", diagnosisApi.NewGetMany())
	router.Put("/diagnosis/{id}", diagnosisApi.NewUpdate())
	router.Delete("/diagnosis/{id}", diagnosisApi.NewDelete())

	router.Post("/treatment", treatmentApi.NewCreate())
	router.Get("/treatment/{id}", treatmentApi.NewGet())
	router.Get("/treatment", treatmentApi.NewGetMany())
	router.Put("/treatment/{id}", treatmentApi.NewUpdate())
	router.Delete("/treatment/{id}", treatmentApi.NewDelete())

	router.Post("/appointment", appointmentApi.NewCreate())
	router.Get("/appointment/{id}", appointmentApi.NewGet())
	router.Get("/appointment", appointmentApi.NewGetMany())
	router.Put("/appointment/{id}", appointmentApi.NewUpdate())
	router.Delete("/appointment/{id}", appointmentApi.NewDelete())

	// init metrics
	router.Handle("/metrics", promhttp.Handler())

	// init server
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}
	done := make(chan os.Signal, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("server started on port 3000")

	<-done

	log.Println("server stopped")
}
