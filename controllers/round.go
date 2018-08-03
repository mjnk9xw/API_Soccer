package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"
	"encoding/json"

	"github.com/astaxie/beego"
)

type RoundController struct {
	beego.Controller
}

// @Title CreateRound
// @Description create round
// @Param	body		body 	models.Round	true		"body for tour content"
// @Success 200 {ResponseRound} ResponseRound
// @Failure 403 body is empty
// @router / [post]
func (r *RoundController) Post() {
	var round models.Round
	TourID := r.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		json.Unmarshal(r.Ctx.Input.RequestBody, &round)
		round.SetRoundID()
		serverstore.Rounds.Set(&round)
		serverstore.RoundIDToListTableID.Set(round.GetRoundID())
		serverstore.TourIDToListRoundID.SetRound(TourID, round.GetRoundID())
	}
	r.Data["json"] = ResponseRound{Data: []models.RoundInfo{&round}, Error: errSC.NewErr(err)}
	r.ServeJSON()
}

// @Title GetAll
// @Description get all rounds
// @Success 200 {ResponseRound} ResponseRound
// @Failure 403 :TourID is empty
// @router / [get]
func (r *RoundController) GetAll() {
	TourID := r.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		rounds, err := ToursByID(TourID)
		r.Data["json"] = ResponseRound{Data: rounds, Error: errSC.NewErr(err)}
	} else {
		r.Data["json"] = ResponseRound{Error: errSC.NewErr(err)}
	}
	r.ServeJSON()
}

// @Title Get
// @Description get round by RoundID
// @Param	RoundID		path 	string	true		"The key for staticblock"
// @Success 200 {ResponseRound} ResponseRound
// @Failure 403 :RoundID is empty
// @router /:RoundID [get]
func (r *RoundController) Get() {
	TourID := r.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := r.Ctx.Input.Param(":RoundID")
		RoundIDInt, err := CheckRound(RoundID)
		if err == errSC.Success {
			round, err := serverstore.Rounds.GetByID(RoundIDInt)
			r.Data["json"] = ResponseRound{Data: []models.RoundInfo{round}, Error: errSC.NewErr(err)}
		} else {
			r.Data["json"] = ResponseRound{Error: errSC.NewErr(err)}
		}
	} else {
		r.Data["json"] = ResponseRound{Error: errSC.NewErr(err)}
	}
	r.ServeJSON()
}

// @Title Update
// @Description update the round
// @Param	body		body 	models.Round	true		"body for round content"
// @Success 200 {ResponseRound} ResponseRound
// @Failure 403 not Round
// @router / [put]
func (r *RoundController) Put() {
	var round models.Round
	TourID := r.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		json.Unmarshal(r.Ctx.Input.RequestBody, &round)
		serverstore.Rounds.Set(&round)
	}
	r.Data["json"] = ResponseRound{Data: []models.RoundInfo{&round}, Error: errSC.NewErr(err)}
	r.ServeJSON()
}

// @Title Delete
// @Description delete the round
// @Param	RoundID		path 	string	true		"The RoundID you want to delete"
// @Success 200 {ResponseRound} ResponseRound
// @Failure 403 RoundID is empty
// @router /:RoundID [delete]
func (r *RoundController) Delete() {
	TourID := r.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := r.Ctx.Input.Param(":RoundID")
		RoundIDInt, err := CheckRound(RoundID)
		if err == errSC.Success {
			serverstore.Rounds.Remove(RoundIDInt)
			serverstore.RoundIDToListTableID.Remove(RoundIDInt)
			serverstore.TourIDToListRoundID.RemoveRound(TourID, RoundIDInt)
		}
	}
	r.Data["json"] = ResponseRound{Error: errSC.NewErr(err)}
	r.ServeJSON()
}
