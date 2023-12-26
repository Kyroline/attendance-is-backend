package route

import (
	getCompensationStudent "attendance-is/controllers/compensation/pbm/get"
	getCompensationHandler "attendance-is/handlers/compensation/pbm/get"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCompensationRoute(db *gorm.DB, r *gin.Engine) {

	getCompensationRepository := getCompensationStudent.NewCompensationGetRepository(db)
	getCompensationService := getCompensationStudent.NewCompensationGetService(getCompensationRepository)
	getCompensationHandler := getCompensationHandler.NewGetCompensationHandler(getCompensationService)

	// compensationService := service.NewCompensationService(db)
	// compensationHandler := compensation.NewCompensationHandler(compensationService)

	// groupStudent := r.Group("api/classs/:classid/student/:studentid/compensation")
	// groupStudent.GET("", getAllCompensationHandler.GetAllCompensationHandler)
	// groupStudent.GET(":id", getCompensationHandler.GetCompensationHandler)

	groupPBM := r.Group("api/compensation")
	groupPBM.GET("", getCompensationHandler.GetCompensationHandler)
}
