package router

import(
	"gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/storage"
	. "github.com/rohit123sinha456/payredapp/handlers"
	. "github.com/rohit123sinha456/payredapp/helpers"
	. "github.com/rohit123sinha456/payredapp/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
	server.router.Use(cors.Default())
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

	//Client Login Routes
	server.router.GET("/google_callback", googleOauthHandler.GoogleCallback)
	server.router.GET("/google_login", googleOauthHandler.GoogleLogin)
	server.router.POST("/clientlogin", clientloginHandler.Login)

	// User Login
	server.router.POST("/userlogin", userloginHandler.Login)

	// User routes
	userprotected := server.router.Group("/api/web", AuthMiddleware())
	userprotected.POST("/users", userHandler.CreateUser)
	userprotected.GET("/users/:id", userHandler.GetUser)
	userprotected.GET("/users", userHandler.GetAllUsers)
	userprotected.PUT("/users/:id", userHandler.UpdateUser)
	userprotected.DELETE("/users/:id", userHandler.DeleteUser)
	

	// Payment routes
	userprotected.POST("/payments", paymentHandler.CreatePayment)
	userprotected.GET("/payments/:id", paymentHandler.GetPayment)
	userprotected.GET("/payments", paymentHandler.GetAllPayments)
	userprotected.PUT("/payments/:id", paymentHandler.UpdatePayment)
	userprotected.DELETE("/payments/:id", paymentHandler.DeletePayment)

	// Quotation routes
    userprotected.POST("/quotations", quotationHandler.CreateQuotation)
    userprotected.GET("/quotations/:id", quotationHandler.GetQuotation)
    userprotected.GET("/quotations", quotationHandler.GetAllQuotations)
    userprotected.PUT("/quotations/:id", quotationHandler.UpdateQuotation)
    userprotected.DELETE("/quotations/:id", quotationHandler.DeleteQuotation)

    // ServiceReport routes
    userprotected.POST("/servicereports", serviceReportHandler.CreateServiceReport)
    userprotected.GET("/servicereports/:id", serviceReportHandler.GetServiceReport)
    userprotected.GET("/servicereports", serviceReportHandler.GetAllServiceReports)
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