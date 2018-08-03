package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"

	"github.com/astaxie/beego"
)

// Operations about PlayerInTour
type PlayersInTourController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all players
// @Success 200 {ResponsePlayer} ResponsePlayer
// @Failure 403 :TourID is empty
// @router / [get]
func (t *PlayersInTourController) GetAll() {
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		players, err := PlayersInTour(TourID)
		t.Data["json"] = ResponsePlayer{Data: players, Error: errSC.NewErr(err)}
	} else {
		t.Data["json"] = ResponseTour{Error: errSC.NewErr(err)}
	}
	t.ServeJSON()
}

// @Title Get
// @Description get player by PlayerID
// @Param	PlayerID		path 	string	true		"The key for staticblock"
// @Success 200 {ResponsePlayer} ResponsePlayer
// @Failure 403 :PlayerID is empty
// @router /:PlayerID [get]
func (t *PlayersInTourController) Get() {
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		PlayerID := t.Ctx.Input.Param(":PlayerID")
		PlayerIDInt, err := CheckPlayer(PlayerID)
		if err == errSC.Success {
			player, err := serverstore.Players.GetByID(PlayerIDInt)
			t.Data["json"] = ResponsePlayer{Data: []models.PlayerInfo{player}, Error: errSC.NewErr(err)}
		} else {
			t.Data["json"] = ResponsePlayer{Error: errSC.NewErr(err)}
		}
	} else {
		t.Data["json"] = ResponsePlayer{Error: errSC.NewErr(err)}
	}
	t.ServeJSON()
}
