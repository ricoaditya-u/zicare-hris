package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ricoaditya-u/hris-master/controllers"
)

func InitializeRoutesMpp(g *gin.RouterGroup) {
	{
		// HEAD DIV/DEPT
		// MPP
		g.GET("/index/:employeeid/", controllers.MppIndex)
		g.POST("/createMpp", controllers.MppCreate)
		g.PUT("/updateMpp/:id", controllers.MppUpdate)
		g.GET("/showMpp/:id", controllers.MppShow)
		// REQ HEADCOUNT
		g.GET("/index/:employeeid/:period", controllers.ListMpp)
		g.GET("/formHeadcount/:mppid", controllers.FormHeadcount) // Form Create
		g.POST("/createHeadcount", controllers.CreateHeadcount)
		g.GET("/showAllHeadcount", controllers.ShowAllHeadcount)

		// HR
		g.GET("/listUnapproveMpp", controllers.MppListUnapprove)
		g.PUT("/approveMpp/:id", controllers.MppApprove)
		g.PUT("/revisionMpp/:id", controllers.MppRevision)
	}
}
