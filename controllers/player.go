package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"
	"encoding/json"

	"github.com/astaxie/beego"
)

type PlayerController struct {
	beego.Controller
}

// @Title CreatePlayer
// @Description create player
// @Param	body		body 	models.Player	true		"body for player content"
// @Success 200 {ResponsePlayer} ResponsePlayer
// @Failure 403 body is empty
// @router / [post]
func (p *PlayerController) Post() {
	var player models.Player
	TourID := p.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := p.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := p.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				json.Unmarshal(p.Ctx.Input.RequestBody, &player)
				player.SetPlayerID()
				serverstore.Players.Set(&player)
				serverstore.TourIDToListPlayerID.SetPlayer(TourID, player.GetPlayerID())
				serverstore.TableIDToListPlayerID.SetPlayer(TableIDInt, player.GetPlayerID())

			}
		}
	}
	p.Data["json"] = ResponsePlayer{Data: []models.PlayerInfo{&player}, Error: errSC.NewErr(err)}
	p.ServeJSON()
}

// @Title Update
// @Description update the player
// @Param	body		body 	models.Player	true		"body for player content"
// @Success 200 {ResponsePlayer} ResponsePlayer
// @Failure 403 not Player
// @router / [put]
func (p *PlayerController) Put() {
	var player models.Player
	TourID := p.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := p.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := p.Ctx.Input.Param(":TableID")
			_, err := CheckTable(TableID)
			if err == errSC.Success {
				json.Unmarshal(p.Ctx.Input.RequestBody, &player)
				serverstore.Players.Set(&player)
			}
		}
	}
	p.Data["json"] = ResponsePlayer{Data: []models.PlayerInfo{&player}, Error: errSC.NewErr(err)}
	p.ServeJSON()
}

// @Title Delete
// @Description delete the player
// @Param	PlayerID		path 	string	true		"The PlayerID you want to delete"
// @Success 200 {ResponsePlayer} ResponsePlayer
// @Failure 403 PlayerID is empty
// @router /:PlayerID [delete]
func (p *PlayerController) Delete() {
	TourID := p.Ctx.Input.Param(":TourID")
	PlayerID := p.Ctx.Input.Param(":PlayerID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := p.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := p.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				PlayerIDInt, err := CheckPlayer(PlayerID)
				if err == errSC.Success {
					serverstore.Players.Remove(PlayerIDInt)
					serverstore.TourIDToListPlayerID.RemovePlayer(TourID, PlayerIDInt)
					serverstore.TableIDToListPlayerID.RemovePlayer(TableIDInt, PlayerIDInt)
				}
			}
		}
	}
	p.Data["json"] = ResponsePlayer{Error: errSC.NewErr(err)}
	p.ServeJSON()
}
