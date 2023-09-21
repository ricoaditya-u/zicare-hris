package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ricoaditya-u/hris-master/controllers"
)

func InitializeRoutesMpp(g *gin.RouterGroup) {
	{
		// HEAD DIV/DEPT
		// MPP
		g.GET("/index", controllers.MppIndex)
		g.POST("/createMpp", controllers.MppCreate)
		g.PUT("/updateMpp/:id", controllers.MppUpdate)
		g.GET("/showMpp/:id", controllers.MppShow)

		// REQ HEADCOUNT
		g.POST("/createHeadcount", controllers.HeadcountCreate)

		// HR
		g.GET("/listUnapproveMpp", controllers.MppListUnapprove)
		g.PUT("/approveMpp/:id", controllers.MppApprove)
		g.PUT("/revisionMpp/:id", controllers.MppRevision)
	}
}
