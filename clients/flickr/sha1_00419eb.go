// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_sha1_00419eb_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("RemotePhoto",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Title", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Url", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Geoposition", types.MakeType(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 0), false},
				types.Field{"Sizes", types.MakeCompoundType(types.MapKind, types.MakeType(ref.Ref{}, 1), types.MakePrimitiveType(types.StringKind)), false},
				types.Field{"Tags", types.MakeCompoundType(types.SetKind, types.MakePrimitiveType(types.StringKind)), false},
			},
			types.Choices{},
		),
		types.MakeStructType("Size",
			[]types.Field{
				types.Field{"Width", types.MakePrimitiveType(types.UInt32Kind), false},
				types.Field{"Height", types.MakePrimitiveType(types.UInt32Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"),
	})
	__mainPackageInFile_sha1_00419eb_CachedRef = types.RegisterPackage(&p)
}

// RemotePhoto

type RemotePhoto struct {
	_Id          string
	_Title       string
	_Url         string
	_Geoposition Geoposition
	_Sizes       MapOfSizeToString
	_Tags        SetOfString

	ref *ref.Ref
}

func NewRemotePhoto() RemotePhoto {
	return RemotePhoto{
		_Id:          "",
		_Title:       "",
		_Url:         "",
		_Geoposition: NewGeoposition(),
		_Sizes:       NewMapOfSizeToString(),
		_Tags:        NewSetOfString(),

		ref: &ref.Ref{},
	}
}

type RemotePhotoDef struct {
	Id          string
	Title       string
	Url         string
	Geoposition GeopositionDef
	Sizes       MapOfSizeToStringDef
	Tags        SetOfStringDef
}

func (def RemotePhotoDef) New() RemotePhoto {
	return RemotePhoto{
		_Id:          def.Id,
		_Title:       def.Title,
		_Url:         def.Url,
		_Geoposition: def.Geoposition.New(),
		_Sizes:       def.Sizes.New(),
		_Tags:        def.Tags.New(),
		ref:          &ref.Ref{},
	}
}

func (s RemotePhoto) Def() (d RemotePhotoDef) {
	d.Id = s._Id
	d.Title = s._Title
	d.Url = s._Url
	d.Geoposition = s._Geoposition.Def()
	d.Sizes = s._Sizes.Def()
	d.Tags = s._Tags.Def()
	return
}

var __typeForRemotePhoto types.Type

func (m RemotePhoto) Type() types.Type {
	return __typeForRemotePhoto
}

func init() {
	__typeForRemotePhoto = types.MakeType(__mainPackageInFile_sha1_00419eb_CachedRef, 0)
	types.RegisterStruct(__typeForRemotePhoto, builderForRemotePhoto, readerForRemotePhoto)
}

func builderForRemotePhoto(values []types.Value) types.Value {
	i := 0
	s := RemotePhoto{ref: &ref.Ref{}}
	s._Id = values[i].(types.String).String()
	i++
	s._Title = values[i].(types.String).String()
	i++
	s._Url = values[i].(types.String).String()
	i++
	s._Geoposition = values[i].(Geoposition)
	i++
	s._Sizes = values[i].(MapOfSizeToString)
	i++
	s._Tags = values[i].(SetOfString)
	i++
	return s
}

func readerForRemotePhoto(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(RemotePhoto)
	values = append(values, types.NewString(s._Id))
	values = append(values, types.NewString(s._Title))
	values = append(values, types.NewString(s._Url))
	values = append(values, s._Geoposition)
	values = append(values, s._Sizes)
	values = append(values, s._Tags)
	return values
}

func (s RemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeForRemotePhoto.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s RemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s RemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForRemotePhoto.Chunks()...)
	chunks = append(chunks, s._Geoposition.Chunks()...)
	chunks = append(chunks, s._Sizes.Chunks()...)
	chunks = append(chunks, s._Tags.Chunks()...)
	return
}

func (s RemotePhoto) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._Id))
	ret = append(ret, types.NewString(s._Title))
	ret = append(ret, types.NewString(s._Url))
	ret = append(ret, s._Geoposition)
	ret = append(ret, s._Sizes)
	ret = append(ret, s._Tags)
	return
}

func (s RemotePhoto) Id() string {
	return s._Id
}

func (s RemotePhoto) SetId(val string) RemotePhoto {
	s._Id = val
	s.ref = &ref.Ref{}
	return s
}

func (s RemotePhoto) Title() string {
	return s._Title
}

func (s RemotePhoto) SetTitle(val string) RemotePhoto {
	s._Title = val
	s.ref = &ref.Ref{}
	return s
}

func (s RemotePhoto) Url() string {
	return s._Url
}

func (s RemotePhoto) SetUrl(val string) RemotePhoto {
	s._Url = val
	s.ref = &ref.Ref{}
	return s
}

func (s RemotePhoto) Geoposition() Geoposition {
	return s._Geoposition
}

func (s RemotePhoto) SetGeoposition(val Geoposition) RemotePhoto {
	s._Geoposition = val
	s.ref = &ref.Ref{}
	return s
}

func (s RemotePhoto) Sizes() MapOfSizeToString {
	return s._Sizes
}

func (s RemotePhoto) SetSizes(val MapOfSizeToString) RemotePhoto {
	s._Sizes = val
	s.ref = &ref.Ref{}
	return s
}

func (s RemotePhoto) Tags() SetOfString {
	return s._Tags
}

func (s RemotePhoto) SetTags(val SetOfString) RemotePhoto {
	s._Tags = val
	s.ref = &ref.Ref{}
	return s
}

// Size

type Size struct {
	_Width  uint32
	_Height uint32

	ref *ref.Ref
}

func NewSize() Size {
	return Size{
		_Width:  uint32(0),
		_Height: uint32(0),

		ref: &ref.Ref{},
	}
}

type SizeDef struct {
	Width  uint32
	Height uint32
}

func (def SizeDef) New() Size {
	return Size{
		_Width:  def.Width,
		_Height: def.Height,
		ref:     &ref.Ref{},
	}
}

func (s Size) Def() (d SizeDef) {
	d.Width = s._Width
	d.Height = s._Height
	return
}

var __typeForSize types.Type

func (m Size) Type() types.Type {
	return __typeForSize
}

func init() {
	__typeForSize = types.MakeType(__mainPackageInFile_sha1_00419eb_CachedRef, 1)
	types.RegisterStruct(__typeForSize, builderForSize, readerForSize)
}

func builderForSize(values []types.Value) types.Value {
	i := 0
	s := Size{ref: &ref.Ref{}}
	s._Width = uint32(values[i].(types.UInt32))
	i++
	s._Height = uint32(values[i].(types.UInt32))
	i++
	return s
}

func readerForSize(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Size)
	values = append(values, types.UInt32(s._Width))
	values = append(values, types.UInt32(s._Height))
	return values
}

func (s Size) Equals(other types.Value) bool {
	return other != nil && __typeForSize.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Size) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Size) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForSize.Chunks()...)
	return
}

func (s Size) ChildValues() (ret []types.Value) {
	ret = append(ret, types.UInt32(s._Width))
	ret = append(ret, types.UInt32(s._Height))
	return
}

func (s Size) Width() uint32 {
	return s._Width
}

func (s Size) SetWidth(val uint32) Size {
	s._Width = val
	s.ref = &ref.Ref{}
	return s
}

func (s Size) Height() uint32 {
	return s._Height
}

func (s Size) SetHeight(val uint32) Size {
	s._Height = val
	s.ref = &ref.Ref{}
	return s
}

// MapOfSizeToString

type MapOfSizeToString struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfSizeToString() MapOfSizeToString {
	return MapOfSizeToString{types.NewTypedMap(__typeForMapOfSizeToString), &ref.Ref{}}
}

type MapOfSizeToStringDef map[SizeDef]string

func (def MapOfSizeToStringDef) New() MapOfSizeToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, k.New(), types.NewString(v))
	}
	return MapOfSizeToString{types.NewTypedMap(__typeForMapOfSizeToString, kv...), &ref.Ref{}}
}

func (m MapOfSizeToString) Def() MapOfSizeToStringDef {
	def := make(map[SizeDef]string)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(Size).Def()] = v.(types.String).String()
		return false
	})
	return def
}

func (m MapOfSizeToString) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfSizeToString.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfSizeToString) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfSizeToString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfSizeToString) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfSizeToString.
var __typeForMapOfSizeToString types.Type

func (m MapOfSizeToString) Type() types.Type {
	return __typeForMapOfSizeToString
}

func init() {
	__typeForMapOfSizeToString = types.MakeCompoundType(types.MapKind, types.MakeType(__mainPackageInFile_sha1_00419eb_CachedRef, 1), types.MakePrimitiveType(types.StringKind))
	types.RegisterValue(__typeForMapOfSizeToString, builderForMapOfSizeToString, readerForMapOfSizeToString)
}

func builderForMapOfSizeToString(v types.Value) types.Value {
	return MapOfSizeToString{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfSizeToString(v types.Value) types.Value {
	return v.(MapOfSizeToString).m
}

func (m MapOfSizeToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfSizeToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfSizeToString) Has(p Size) bool {
	return m.m.Has(p)
}

func (m MapOfSizeToString) Get(p Size) string {
	return m.m.Get(p).(types.String).String()
}

func (m MapOfSizeToString) MaybeGet(p Size) (string, bool) {
	v, ok := m.m.MaybeGet(p)
	if !ok {
		return "", false
	}
	return v.(types.String).String(), ok
}

func (m MapOfSizeToString) Set(k Size, v string) MapOfSizeToString {
	return MapOfSizeToString{m.m.Set(k, types.NewString(v)), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfSizeToString) Remove(p Size) MapOfSizeToString {
	return MapOfSizeToString{m.m.Remove(p), &ref.Ref{}}
}

type MapOfSizeToStringIterCallback func(k Size, v string) (stop bool)

func (m MapOfSizeToString) Iter(cb MapOfSizeToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(Size), v.(types.String).String())
	})
}

type MapOfSizeToStringIterAllCallback func(k Size, v string)

func (m MapOfSizeToString) IterAll(cb MapOfSizeToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(Size), v.(types.String).String())
	})
}

func (m MapOfSizeToString) IterAllP(concurrency int, cb MapOfSizeToStringIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(Size), v.(types.String).String())
	})
}

type MapOfSizeToStringFilterCallback func(k Size, v string) (keep bool)

func (m MapOfSizeToString) Filter(cb MapOfSizeToStringFilterCallback) MapOfSizeToString {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(Size), v.(types.String).String())
	})
	return MapOfSizeToString{out, &ref.Ref{}}
}

// SetOfString

type SetOfString struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfString() SetOfString {
	return SetOfString{types.NewTypedSet(__typeForSetOfString), &ref.Ref{}}
}

type SetOfStringDef map[string]bool

func (def SetOfStringDef) New() SetOfString {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.NewString(d)
		i++
	}
	return SetOfString{types.NewTypedSet(__typeForSetOfString, l...), &ref.Ref{}}
}

func (s SetOfString) Def() SetOfStringDef {
	def := make(map[string]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(types.String).String()] = true
		return false
	})
	return def
}

func (s SetOfString) Equals(other types.Value) bool {
	return other != nil && __typeForSetOfString.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SetOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.Type().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

func (s SetOfString) ChildValues() []types.Value {
	return append([]types.Value{}, s.s.ChildValues()...)
}

// A Noms Value that describes SetOfString.
var __typeForSetOfString types.Type

func (m SetOfString) Type() types.Type {
	return __typeForSetOfString
}

func init() {
	__typeForSetOfString = types.MakeCompoundType(types.SetKind, types.MakePrimitiveType(types.StringKind))
	types.RegisterValue(__typeForSetOfString, builderForSetOfString, readerForSetOfString)
}

func builderForSetOfString(v types.Value) types.Value {
	return SetOfString{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfString(v types.Value) types.Value {
	return v.(SetOfString).s
}

func (s SetOfString) Empty() bool {
	return s.s.Empty()
}

func (s SetOfString) Len() uint64 {
	return s.s.Len()
}

func (s SetOfString) Has(p string) bool {
	return s.s.Has(types.NewString(p))
}

type SetOfStringIterCallback func(p string) (stop bool)

func (s SetOfString) Iter(cb SetOfStringIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(types.String).String())
	})
}

type SetOfStringIterAllCallback func(p string)

func (s SetOfString) IterAll(cb SetOfStringIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(types.String).String())
	})
}

func (s SetOfString) IterAllP(concurrency int, cb SetOfStringIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(types.String).String())
	})
}

type SetOfStringFilterCallback func(p string) (keep bool)

func (s SetOfString) Filter(cb SetOfStringFilterCallback) SetOfString {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(types.String).String())
	})
	return SetOfString{out, &ref.Ref{}}
}

func (s SetOfString) Insert(p ...string) SetOfString {
	return SetOfString{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfString) Remove(p ...string) SetOfString {
	return SetOfString{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfString) Union(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfString) Subtract(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfString) Any() string {
	return s.s.Any().(types.String).String()
}

func (s SetOfString) fromStructSlice(p []SetOfString) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfString) fromElemSlice(p []string) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.NewString(v)
	}
	return r
}
