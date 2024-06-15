package main

import (
	"fmt"
	application "golang-boilerplate"
	"golang-boilerplate/config"
	"golang-boilerplate/docs"
)

//	@title			Golang boilerplate
//	@version		1.0
//	@description	A golang boilerplate

//	@contact.name	Phuckhoa
//	@contact.url	https://www.facebook.com/profile.php?id=61560223583957
//	@contact.email	phuckhoa81@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8888
// @BasePath	/api/v1
func main() {
	config := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.HTTP.AppHost, config.HTTP.AppPort)

	application.Start(config)
}
