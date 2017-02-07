// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"WebService/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	//Incluyendo el CORS
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_ingreso",
			beego.NSInclude(
				&controllers.EstadoIngresoController{},
			),
		),

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/forma_ingreso",
			beego.NSInclude(
				&controllers.FormaIngresoController{},
			),
		),

		beego.NSNamespace("/ingreso_concepto",
			beego.NSInclude(
				&controllers.IngresoConceptoController{},
			),
		),

		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/ingreso",
			beego.NSInclude(
				&controllers.IngresoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
