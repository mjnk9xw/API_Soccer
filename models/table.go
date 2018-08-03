package models

type Table struct {
	TableID   int64        `json:"id"`
	TableName string       `json:"name"`
	Players   []PlayerInfo `json:"players"`
	Matchs    []MatchInfo  `json:"matchs"`
}

func (t *Table) SetTableID() {
	t.TableID = genNextTableID()
}
func (t *Table) GetTableID() int64 {
	return t.TableID
}
func (t *Table) SetPlayers(p []PlayerInfo) {
	t.Players = make([]PlayerInfo, 0)
	t.Players = append(t.Players, p...)
}
func (t *Table) SetMatchs(m []MatchInfo) {
	t.Matchs = make([]MatchInfo, 0)
	t.Matchs = append(t.Matchs, m...)
}
