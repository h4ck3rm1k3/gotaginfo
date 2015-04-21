package model
import "github.com/h4ck3rm1k3/taginfo/database/taginfo"

func Set(src * taginfo_model.Keys, dest * Keys){
	dest.SetKey(src.Key)
	dest.SetCountall(src.Countall)
	dest.SetCountnodes(src.Countnodes)
	dest.SetCountways(src.Countways)
	dest.SetCountrelations(src.Countrelations)
	dest.SetValuesall(src.Valuesall)
	dest.SetValuesnodes(src.Valuesnodes)
	dest.SetValuesways(src.Valuesways)
	dest.SetValuesrelations(src.Valuesrelations)
	dest.SetUsersall(src.Usersall)
	dest.SetUsersnodes(src.Usersnodes)
	dest.SetUsersways(src.Usersways)
	dest.SetUsersrelations(src.Usersrelations)
	dest.SetCellsnodes(src.Cellsnodes)
	dest.SetCellsways(src.Cellsways)
	dest.SetInwiki(src.Inwiki)
	dest.SetInprojects(src.Inprojects)
	dest.SetCharacters(src.Characters)
}

func Get(src * Keys, dest * taginfo_model.Keys){
	dest.Key=src.Key()
	dest.Countall=src.Countall()
	dest.Countnodes=src.Countnodes()
	dest.Countways=src.Countways()
	dest.Countrelations=src.Countrelations()
	dest.Valuesall=src.Valuesall()
	dest.Valuesnodes=src.Valuesnodes()
	dest.Valuesways=src.Valuesways()
	dest.Valuesrelations=src.Valuesrelations()
	dest.Usersall=src.Usersall()
	dest.Usersnodes=src.Usersnodes()
	dest.Usersways=src.Usersways()
	dest.Usersrelations=src.Usersrelations()
	dest.Cellsnodes=src.Cellsnodes()
	dest.Cellsways=src.Cellsways()
	dest.Inwiki=src.Inwiki()
	dest.Inprojects=src.Inprojects()
	dest.Characters=src.Characters()
}
