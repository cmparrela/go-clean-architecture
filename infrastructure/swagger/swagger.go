package swagger

import "github.com/gofiber/swagger"

func InitSwagger() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		URL:          "../../../docs/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	})
}
