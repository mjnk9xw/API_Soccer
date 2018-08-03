package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:MatchID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:MatchController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:MatchID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:PlayerID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:PlayersInTourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:PlayersInTourController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:PlayersInTourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:PlayersInTourController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:PlayerID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:RoundID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:RoundController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:RoundID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:TableID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:TableID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:TourID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"] = append(beego.GlobalControllerRouter["SoccerBeego/controllers:TourController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:TourID`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
