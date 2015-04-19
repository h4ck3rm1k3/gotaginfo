package main

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
)

type Keys C.Struct

func NewKeys(s *C.Segment) Keys           { return Keys(s.NewStruct(128, 2)) }
func NewRootKeys(s *C.Segment) Keys       { return Keys(s.NewRootStruct(128, 2)) }
func AutoNewKeys(s *C.Segment) Keys       { return Keys(s.NewStructAR(128, 2)) }
func ReadRootKeys(s *C.Segment) Keys      { return Keys(s.Root(0).ToStruct()) }
func (s Keys) Key() string                { return C.Struct(s).GetObject(0).ToText() }
func (s Keys) SetKey(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Keys) Countall() int64            { return int64(C.Struct(s).Get64(0)) }
func (s Keys) SetCountall(v int64)        { C.Struct(s).Set64(0, uint64(v)) }
func (s Keys) Countnodes() int64          { return int64(C.Struct(s).Get64(8)) }
func (s Keys) SetCountnodes(v int64)      { C.Struct(s).Set64(8, uint64(v)) }
func (s Keys) Countways() int64           { return int64(C.Struct(s).Get64(16)) }
func (s Keys) SetCountways(v int64)       { C.Struct(s).Set64(16, uint64(v)) }
func (s Keys) Countrelations() int64      { return int64(C.Struct(s).Get64(24)) }
func (s Keys) SetCountrelations(v int64)  { C.Struct(s).Set64(24, uint64(v)) }
func (s Keys) Valuesall() int64           { return int64(C.Struct(s).Get64(32)) }
func (s Keys) SetValuesall(v int64)       { C.Struct(s).Set64(32, uint64(v)) }
func (s Keys) Valuesnodes() int64         { return int64(C.Struct(s).Get64(40)) }
func (s Keys) SetValuesnodes(v int64)     { C.Struct(s).Set64(40, uint64(v)) }
func (s Keys) Valuesways() int64          { return int64(C.Struct(s).Get64(48)) }
func (s Keys) SetValuesways(v int64)      { C.Struct(s).Set64(48, uint64(v)) }
func (s Keys) Valuesrelations() int64     { return int64(C.Struct(s).Get64(56)) }
func (s Keys) SetValuesrelations(v int64) { C.Struct(s).Set64(56, uint64(v)) }
func (s Keys) Usersall() int64            { return int64(C.Struct(s).Get64(64)) }
func (s Keys) SetUsersall(v int64)        { C.Struct(s).Set64(64, uint64(v)) }
func (s Keys) Usersnodes() int64          { return int64(C.Struct(s).Get64(72)) }
func (s Keys) SetUsersnodes(v int64)      { C.Struct(s).Set64(72, uint64(v)) }
func (s Keys) Usersways() int64           { return int64(C.Struct(s).Get64(80)) }
func (s Keys) SetUsersways(v int64)       { C.Struct(s).Set64(80, uint64(v)) }
func (s Keys) Usersrelations() int64      { return int64(C.Struct(s).Get64(88)) }
func (s Keys) SetUsersrelations(v int64)  { C.Struct(s).Set64(88, uint64(v)) }
func (s Keys) Cellsnodes() int64          { return int64(C.Struct(s).Get64(96)) }
func (s Keys) SetCellsnodes(v int64)      { C.Struct(s).Set64(96, uint64(v)) }
func (s Keys) Cellsways() int64           { return int64(C.Struct(s).Get64(104)) }
func (s Keys) SetCellsways(v int64)       { C.Struct(s).Set64(104, uint64(v)) }
func (s Keys) Inwiki() int64              { return int64(C.Struct(s).Get64(112)) }
func (s Keys) SetInwiki(v int64)          { C.Struct(s).Set64(112, uint64(v)) }
func (s Keys) Inprojects() int64          { return int64(C.Struct(s).Get64(120)) }
func (s Keys) SetInprojects(v int64)      { C.Struct(s).Set64(120, uint64(v)) }
func (s Keys) Characters() string         { return C.Struct(s).GetObject(1).ToText() }
func (s Keys) SetCharacters(v string)     { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Keys) MarshalJSON() (bs []byte, err error) { return }

type Keys_List C.PointerList

func NewKeysList(s *C.Segment, sz int) Keys_List { return Keys_List(s.NewCompositeList(128, 2, sz)) }
func (s Keys_List) Len() int                     { return C.PointerList(s).Len() }
func (s Keys_List) At(i int) Keys                { return Keys(C.PointerList(s).At(i).ToStruct()) }
func (s Keys_List) ToArray() []Keys {
	n := s.Len()
	a := make([]Keys, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Keys_List) Set(i int, item Keys) { C.PointerList(s).Set(i, C.Object(item)) }
