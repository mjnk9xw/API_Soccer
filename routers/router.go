package routers

import (
	"SoccerBeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// @APIVersion 1.10.1
	// @Title beego TestBeego API
	// @Description beego has a very cool tools to autogenerate documents for your API
	// @Contact minhnguyen998vp@gmail.com
	// @TermsOfServiceUrl http://beego.me/
	// @License Apache 2.0
	// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/tour",
			beego.NSInclude(
				&controllers.TourController{},
			),
		),
		beego.NSNamespace("/tour/:TourID/round",
			beego.NSInclude(
				&controllers.RoundController{},
			),
		),
		beego.NSNamespace("/tour/:TourID/round/:RoundID/table",
			beego.NSInclude(
				&controllers.TableController{},
			),
		),
		beego.NSNamespace("/tour/:TourID/round/:RoundID/table/:TableID/player",
			beego.NSInclude(
				&controllers.PlayerController{},
			),
		),
		beego.NSNamespace("/tour/:TourID/round/:RoundID/table/:TableID/match",
			beego.NSInclude(
				&controllers.MatchController{},
			),
		),
		beego.NSNamespace("/tour/:TourID/player",
			beego.NSInclude(
				&controllers.PlayersInTourController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
