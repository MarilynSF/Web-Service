package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["WebService/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["WebService/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:IngresoController"] = append(beego.GlobalControllerRouter["WebService/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["WebService/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
