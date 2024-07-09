package router

import(
	"gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/storage"
	. "github.com/rohit123sinha456/payredapp/handlers"
	. "github.com/rohit123sinha456/payredapp/helpers"
	. "github.com/rohit123sinha456/payredapp/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

func (server *APIServer) Init() {
	server.router.Use(cors.Default())
	userHandler := &UserHandler{Repository: &server.UserStore}
    paymentHandler := &PaymentHandler{Repository: &server.PaymentStore}
	quotationHandler := &QuotationHandler{Repository: &server.QuotationStore}
    serviceReportHandler := &ServiceReportHandler{Repository: &server.ServiceReportStore}
    nocHandler := &NOCHandler{Repository: &server.NocStore}
    clientHandler := &ClientHandler{Repository: &server.ClientStore}
	loginHandler := &LoginHandler{UserRepo: &server.ClientStore}

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
	server.router.POST("/login", loginHandler.Login)

	// User routes
	server.router.POST("/users", userHandler.CreateUser)
	server.router.GET("/users/:id", userHandler.GetUser)
	server.router.GET("/users", userHandler.GetAllUsers)
	server.router.PUT("/users/:id", userHandler.UpdateUser)
	server.router.DELETE("/users/:id", userHandler.DeleteUser)
	

	// Payment routes
	server.router.POST("/payments", paymentHandler.CreatePayment)
	server.router.GET("/payments/:id", paymentHandler.GetPayment)
	server.router.GET("/payments", paymentHandler.GetAllPayments)
	server.router.PUT("/payments/:id", paymentHandler.UpdatePayment)
	server.router.DELETE("/payments/:id", paymentHandler.DeletePayment)

	// Quotation routes
    server.router.POST("/quotations", quotationHandler.CreateQuotation)
    server.router.GET("/quotations/:id", quotationHandler.GetQuotation)
    server.router.GET("/quotations", quotationHandler.GetAllQuotations)
    server.router.PUT("/quotations/:id", quotationHandler.UpdateQuotation)
    server.router.DELETE("/quotations/:id", quotationHandler.DeleteQuotation)

    // ServiceReport routes
    server.router.POST("/servicereports", serviceReportHandler.CreateServiceReport)
    server.router.GET("/servicereports/:id", serviceReportHandler.GetServiceReport)
    server.router.GET("/servicereports", serviceReportHandler.GetAllServiceReports)
    server.router.PUT("/servicereports/:id", serviceReportHandler.UpdateServiceReport)
    server.router.DELETE("/servicereports/:id", serviceReportHandler.DeleteServiceReport)

    // NOC routes
    server.router.POST("/nocs", nocHandler.CreateNOC)
    server.router.GET("/nocs/:id", nocHandler.GetNOC)
    server.router.GET("/nocs", nocHandler.GetAllNOCs)
    server.router.PUT("/nocs/:id", nocHandler.UpdateNOC)
    server.router.DELETE("/nocs/:id", nocHandler.DeleteNOC)

    // Client routes
    server.router.POST("/clients", clientHandler.CreateClient)
    server.router.GET("/clients/:id", clientHandler.GetClient)
    server.router.GET("/clients", clientHandler.GetAllClients)
    server.router.PUT("/clients/:id", clientHandler.UpdateClient)
    server.router.DELETE("/clients/:id", clientHandler.DeleteClient)
	server.router.GET("/clients/email", clientHandler.GetClientByEmail)

	// Email reminder route
	server.router.POST("/send-reminders", emailHandler.SendReminders)



	// API For app to consume
	protected := server.router.Group("/api/app", AuthMiddleware())
	protected.GET("/payments/reminders", paymentHandler.GetPaymentReminders)
	protected.GET("/quotations/reminders", quotationHandler.GetQuotationReminders)
	protected.GET("/service-reports", serviceReportHandler.GetServiceReports)
	protected.GET("/nocs", nocHandler.GetNOCs)
}

func (server *APIServer) Run() {
	server.Init()
	server.router.Run(server.listenAddr)
}