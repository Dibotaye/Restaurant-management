package routes

import(
controller "golang-Restaurant-management/controllers" 
"github.com/gin-gonic/gin"
)
func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoice", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/invoice_id", controller.updateInvoice())

}