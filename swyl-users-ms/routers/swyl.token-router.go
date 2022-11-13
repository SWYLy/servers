/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package routers

// @imports
import (
	"Swyl/servers/swyl-users-ms/controllers"

	"github.com/gin-gonic/gin"
)

// @notice Method of UserController struct
//
// @dev Declares a list of endpoints
 func TokenRoutes(rg *gin.RouterGroup) {
	rg.POST("/generate-access-token", controllers.GenerateAccessToken)
 }