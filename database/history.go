// GENERATED BY dbtogo (github.com/kdar/dbtogo); DO NOT EDIT
// ---args: /home/h4ck3rm1k3/goroot_sql/bin/kdb sqlite3 taginfo-history.db
package model

type Arger interface {
  Args() []interface{}
}

var InsertStmts = map[string]string{
  "Historystats": "INSERT INTO Historystats (Udate,Key,Value) VALUES (?, ?, ?)",
}
var SelectStmts = map[string]string{
  "Historystats": "SELECT Udate,Key,Value FROM Historystats",
}

type Historystats struct {
  Udate string
  Key   string
  Value string
}

func (t *Historystats) Args() []interface{} { return []interface{}{&t.Udate, &t.Key, &t.Value} }