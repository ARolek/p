//+build proto

package p

//go:proto ignore
type T int

//go:proto F=Builtin,String T=F.lower
func F(v T) *T {
	return &v
}