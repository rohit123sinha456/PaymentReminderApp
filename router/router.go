package router

import(
	"gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/storage"
	. "github.com/rohit123sinha456/payredapp/handlers"
	. "github.com/rohit123sinha456/payredapp/helpers"
	. "github.com/rohit123sinha456/payredapp/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"os"
	// "log/slog"
)

type APIServer struct{
	listenAddr string
	router *gin.Engine
	DB *gorm.DB
	UserStore UserRepository
	PaymentStore PaymentRepository
	QuotationStore QuotationRepository
	ServiceReportStore ServiceReportRepository
	NocStore NOCRepository
	ClientStore ClientRepository


}

func NewAPIServer (listenAddr string, db *gorm.DB) *APIServer {
	apiserver :=  &APIServer{
		listenAddr : listenAddr,
		router : gin.Default(),
		DB : db,
		UserStore : UserRepository{DB: db},
		PaymentStore : PaymentRepository{DB: db},
		QuotationStore : QuotationRepository{DB: db},
		ServiceReportStore : ServiceReportRepository{DB: db},
		NocStore : NOCRepository{DB: db},
		ClientStore : ClientRepository{DB: db},
	}
	return apiserver
}
// func InitializeLogger() {
// 	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic("Failed to open log file")
// 	}

// 	// Create a file handler
// 	fileHandler := handler.NewIOWriterHandler(logFile, slog.LevelDebug )

// 	// Create a console handler
// 	consoleHandler := handler.NewConsoleHandler(slog.LevelDebug )

// 	// Add the handlers to the logger
// 	slog.PushHandler(fileHandler)
// 	slog.PushHandler(consoleHandler)

// 	// Set Gin to use the custom logger
// 	gin.DefaultWriter = slog.Writer()
// }

func (server *APIServer) Init() {
	var userprotected *gin.RouterGroup
	var userprotectedpages *gin.RouterGroup

	server.router.Use(cors.Default())
	server.router.Static("/assets", os.Getenv("TEMPLATEPATH")+"assets")
	server.router.LoadHTMLGlob(os.Getenv("TEMPLATEPATH")+"templates/*")
	// InitializeLogger()
	// Middleware to log all requests
	// server.router.Use(func(c *gin.Context) {
	// 	slog.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
	// 	c.Next()
	// 	slog.Infof("Response status: %d", c.Writer.Status())
	// })

	userHandler := &UserHandler{Repository: &server.UserStore}
    paymentHandler := &PaymentHandler{Repository: &server.PaymentStore}
	quotationHandler := &QuotationHandler{Repository: &server.QuotationStore}
    serviceReportHandler := &ServiceReportHandler{Repository: &server.ServiceReportStore}
    nocHandler := &NOCHandler{Repository: &server.NocStore}
    clientHandler := &ClientHandler{Repository: &server.ClientStore}
	clientloginHandler := &ClientloginHandler{UserRepo: &server.ClientStore}
	userloginHandler := &UserloginHandler{UserRepo: &server.UserStore}

	

	// Email configuration
	emailConfig := GetConfig()
	googleoauthConfig := GoogleConfig()

	// Initialize email handler
	emailHandler := &EmailHandler{DB: server.DB, Config: emailConfig}

	// Initialise Google Oauth Handler
	googleOauthHandler := &Config{GoogleLoginConfig: googleoauthConfig}



	server.router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	//Client Login Routes
	server.router.GET("/google_callback", googleOauthHandler.GoogleCallback)
	server.router.GET("/google_login", googleOauthHandler.GoogleLogin)
	server.router.POST("/clientlogin", clientloginHandler.Login)

	// User Login
	server.router.POST("/userlogin", userloginHandler.Login)

	// For production for token based auth
	if(os.Getenv("MODE") == "dev"){
		//For development for no auth
		userprotected = server.router.Group("/api/web")
	}else{
		userprotected = server.router.Group("/api/web", AuthMiddleware())
		userprotectedpages = server.router.Group("/")
	}

	// Token protected WEB Routes

	userprotectedpages.GET("/clientspage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients.tmpl", gin.H{
			"title": "Main website",
		})
	})

	userprotectedpages.GET("/paymentspage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "payments.tmpl", gin.H{
			"title": "Main website",
		})
	})

	userprotectedpages.GET("/quotationspage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "quotations.tmpl", gin.H{
			"title": "Main website",
		})
	})
	userprotectedpages.GET("/service-reportspage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sreports.tmpl", gin.H{
			"title": "Main website",
		})
	})
	userprotectedpages.GET("/nocspage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nocs.tmpl", gin.H{
			"title": "Main website",
		})
	})



	userprotectedpages.GET("/payments_clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients_payments.tmpl", gin.H{
			"title": "Main website",
		})
	})
	userprotectedpages.GET("/quotations_clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients_quotations.tmpl", gin.H{
			"title": "Main website",
		})
	})
	userprotectedpages.GET("/sreports_clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients_sreports.tmpl", gin.H{
			"title": "Main website",
		})
	})
	userprotectedpages.GET("/nocs_clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients_nocs.tmpl", gin.H{
			"title": "Main website",
		})
	})


	// User routes
	userprotected.POST("/users", userHandler.CreateUser)
	userprotected.GET("/users/:id", userHandler.GetUser)
	userprotected.GET("/users", userHandler.GetAllUsers)
	userprotected.PUT("/users/:id", userHandler.UpdateUser)
	userprotected.DELETE("/users/:id", userHandler.DeleteUser)
	

	// Payment routes
	userprotected.POST("/payments", paymentHandler.CreatePayment)
	userprotected.GET("/payments/:id", paymentHandler.GetPayment)
	userprotected.GET("/payments", paymentHandler.GetAllPayments)
	userprotected.GET("/payments/client/:client_id", paymentHandler.GetPaymentsByClientID) // New route
	userprotected.PUT("/payments/:id", paymentHandler.UpdatePayment)
	userprotected.DELETE("/payments/:id", paymentHandler.DeletePayment)

	// Quotation routes
    userprotected.POST("/quotations", quotationHandler.CreateQuotation)
    userprotected.GET("/quotations/:id", quotationHandler.GetQuotation)
    userprotected.GET("/quotations", quotationHandler.GetAllQuotations)
	userprotected.GET("/quotations/client/:client_id", quotationHandler.GetQuotationByClientID) // New route
    userprotected.PUT("/quotations/:id", quotationHandler.UpdateQuotation)
    userprotected.DELETE("/quotations/:id", quotationHandler.DeleteQuotation)

    // ServiceReport routes
    userprotected.POST("/servicereports", serviceReportHandler.CreateServiceReport)
    userprotected.GET("/servicereports/:id", serviceReportHandler.GetServiceReport)
    userprotected.GET("/servicereports", serviceReportHandler.GetAllServiceReports)
	userprotected.GET("/servicereports/client/:client_id", serviceReportHandler.GetServiceReportsByClientID) // New route
    userprotected.PUT("/servicereports/:id", serviceReportHandler.UpdateServiceReport)
    userprotected.DELETE("/servicereports/:id", serviceReportHandler.DeleteServiceReport)

    // NOC routes
    userprotected.POST("/nocs", nocHandler.CreateNOC)
    userprotected.GET("/nocs/:id", nocHandler.GetNOC)
    userprotected.GET("/nocs", nocHandler.GetAllNOCs)
    userprotected.PUT("/nocs/:id", nocHandler.UpdateNOC)
    userprotected.DELETE("/nocs/:id", nocHandler.DeleteNOC)

    // Client routes
    userprotected.POST("/clients", clientHandler.CreateClient)
    userprotected.GET("/clients/:id", clientHandler.GetClient)
    userprotected.GET("/clients", clientHandler.GetAllClients)
    userprotected.PUT("/clients/:id", clientHandler.UpdateClient)
    userprotected.DELETE("/clients/:id", clientHandler.DeleteClient)
	userprotected.GET("/clients/email", clientHandler.GetClientByEmail)

	// Email reminder route
	userprotected.POST("/send-reminders", emailHandler.SendReminders)
	userprotected.POST("/send-quotation-reminders", emailHandler.SendQuotationReminders)
	userprotected.POST("/send-sreport-reminders", emailHandler.SendServiceReportReminders)





	// API For app to consume
	protected := server.router.Group("/api/app", AuthMiddleware())
	protected.GET("/payments/reminders", paymentHandler.GetPaymentReminders)
	protected.GET("/quotations/reminders", quotationHandler.GetQuotationReminders)
	protected.GET("/service-reports", serviceReportHandler.GetServiceReports)
	protected.GET("/nocs", nocHandler.GetNOCs)
    protected.PUT("/clients/:id", clientHandler.UpdateClient)

}

func (server *APIServer) Run() {
	server.Init()
	server.router.Run(server.listenAddr)
}