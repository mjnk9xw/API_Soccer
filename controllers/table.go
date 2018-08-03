package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"
	"encoding/json"

	"github.com/astaxie/beego"
)

type TableController struct {
	beego.Controller
}

// @Title CreateTable
// @Description create table
// @Param	body		body 	models.Table	true		"body for tour content"
// @Success 200 {ResponseTable} ResponseTable
// @Failure 403 body is empty
// @router / [post]
func (t *TableController) Post() {
	var table models.Table
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := t.Ctx.Input.Param(":RoundID")
		RoundIDInt, err := CheckRound(RoundID)
		if err == errSC.Success {
			json.Unmarshal(t.Ctx.Input.RequestBody, &table)
			table.SetTableID()
			serverstore.Tables.Set(&table)
			serverstore.RoundIDToListTableID.SetTable(RoundIDInt, table.GetTableID())
			serverstore.TableIDToListMatchID.Set(table.GetTableID())
			serverstore.TableIDToListPlayerID.Set(table.GetTableID())
		}
	}
	t.Data["json"] = ResponseTable{Data: []models.TableInfo{&table}, Error: errSC.NewErr(err)}
	t.ServeJSON()
}

// @Title GetAll
// @Description get all tables
// @Success 200 {ResponseTable} ResponseTable
// @Failure 403 :RoundID is empty
// @router / [get]
func (t *TableController) GetAll() {
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := t.Ctx.Input.Param(":RoundID")
		RoundIDInt, err := CheckRound(RoundID)
		if err == errSC.Success {
			tables, err := RoundByID(RoundIDInt)
			t.Data["json"] = ResponseTable{Data: tables, Error: errSC.NewErr(err)}
		} else {
			t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
		}
	} else {
		t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
	}
	t.ServeJSON()
}

// @Title Get
// @Description get tour by TableID
// @Param	TableID		path 	string	true		"The key for staticblock"
// @Success 200 {ResponseTable} ResponseTable
// @Failure 403 :TableID is empty
// @router /:TableID [get]
func (t *TableController) Get() {
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := t.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := t.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				table, err := serverstore.Tables.GetByID(TableIDInt)
				lstPlayers, _ := PlayersInTable(TableIDInt)
				table.SetPlayers(lstPlayers)
				lstMatchs, _ := MatchInTable(TableIDInt)
				table.SetMatchs(lstMatchs)
				t.Data["json"] = ResponseTable{Data: []models.TableInfo{table}, Error: errSC.NewErr(err)}
			} else {
				t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
			}
		} else {
			t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
		}
	} else {
		t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
	}
	t.ServeJSON()
}

// @Title Update
// @Description update the table
// @Param	body		body 	models.Table	true		"body for user content"
// @Success 200 {ResponseTable} ResponseTable
// @Failure 403 not Table
// @router / [put]
func (t *TableController) Put() {
	var table models.Table
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := t.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			json.Unmarshal(t.Ctx.Input.RequestBody, &table)
			serverstore.Tables.Set(&table)
		}
	}
	t.Data["json"] = ResponseTable{Data: []models.TableInfo{&table}, Error: errSC.NewErr(err)}
	t.ServeJSON()
}

// @Title Delete
// @Description delete the table
// @Param	TableID		path 	string	true		"The TableID you want to delete"
// @Success 200 {ResponseTable} ResponseTable
// @Failure 403 TableID is empty
// @router /:TableID [delete]
func (t *TableController) Delete() {
	TourID := t.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := t.Ctx.Input.Param(":RoundID")
		RoundIDInt, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := t.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				serverstore.Tables.Remove(TableIDInt)
				serverstore.RoundIDToListTableID.RemoveTable(RoundIDInt, TableIDInt)
				serverstore.TableIDToListMatchID.Remove(TableIDInt)
				serverstore.TableIDToListPlayerID.Remove(TableIDInt)

			}
		}
	}
	t.Data["json"] = ResponseTable{Error: errSC.NewErr(err)}
	t.ServeJSON()
}
