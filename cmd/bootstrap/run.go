package bootstrap

// @title           Swagger Example API
// @version         2.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @securityDefinitions.apikey	ApiKeyAuth
import (
	"github.com/manot40/gomenu/internal"
)

func Init() {
	internal.Serve()
}
