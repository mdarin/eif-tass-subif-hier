//
// The empty interface type
//
// The empty interface is crucially important for idiomatic Go. Delaying type-checking until
// runtime makes the language feels more dynamic without completely sacrificing strong typing.
package main

import(
	"fmt"
)


// graph
type G struct {
	vertex interface{
		AddVertex(*G,int) error
		RemoveVertex(*G,int) error
		ConnectVertex(*G,int,int,int) error
		DisconnectVertex(*G,int,int) error
		EdgeScanStart(*G,int,*ES) error
		EdgeScanEnd(*ES) error
		EdgeScanNext(*ES) error
	}
}
func NewG() (*G,error) {
	return new(G),nil
}
func (g *G) MakeG(t int) (err error) {
	switch t {
	case 0:
		g.vertex,err = NewAM()
	default:
		g.vertex,err = NewAL()
	}
	return err
}
func (g *G) AddVertex(index int) error {
	return g.vertex.AddVertex(g,index)
}
func (g *G) ConnectVertex(source,destination,cost int) error {
	return g.vertex.ConnectVertex(g,source,destination,cost)
}
func (g *G) DisconnectVertex(source,destination int) error {
	return g.vertex.DisconnectVertex(g,source,destination)
}
func (g *G) FindVertex(tag interface{}) error {
	//switch tag.(*AL)
	//TODO: Type assertion
	return nil
}


// edges
type ES struct {
	g *G
	cost int
	source int
	dest int
	internal interface{}
}
func NewES() (*ES,error) {
	return new(ES),nil
}



// adj matrix
type AM struct {

}
func NewAM() (*AM,error) {
	fmt.Println("AM::New")
	return new(AM),nil
}
func (am *AM) AddVertex(g *G, index int) error {
	fmt.Println("AM::AddVertex")
	return nil
}
func (am *AM) RemoveVertex(g *G, index int) error {
	fmt.Println("AM::RemoveVertex")
	return nil
}
func (am *AM) ConnectVertex(g *G, source, destination, cost int) error {
	fmt.Println("AM::ConnectVertex")
	return nil
}
func (am *AM) DisconnectVertex(g *G, source int, destination int) error {
	fmt.Println("AM::DisconnectVertex")
	return nil
}
func (am *AM) EdgeScanStart(g *G,index int,es *ES) error {
	fmt.Println("AM::EdgeScanStart")
	return nil
}
func (am *AM) EdgeScanEnd(es *ES) error {
	fmt.Println("AM::EdgeScanEnd")
	return nil
}
func (am *AM) EdgeScanNext(es *ES) error {
	fmt.Println("AM::EdgeScanNext")
	return nil
}


// adj list
type AL struct {

}
func NewAL() (*AL,error) {
	fmt.Println("AL::New")
	return new(AL),nil
}
func (al *AL) AddVertex(g *G, index int) error {
	fmt.Println("AL::AddVertex")
	return nil
}
func (al *AL) RemoveVertex(g *G, index int) error {
	fmt.Println("AL::RemoveVertex")
	return nil
}
func (al *AL) ConnectVertex(g *G, surce, destination, cost int) error {
	fmt.Println("AL::ConnectVertex")
	return nil
}
func (al *AL) DisconnectVertex(g *G, source int, destination int) error {
	fmt.Println("AL::DisconnectVertex")
	return nil
}
func (al *AL) EdgeScanStart(g *G,index int,es *ES) error {
	fmt.Println("AL::EdgeScanStart")
	return nil
}
func (al *AL) EdgeScanEnd(es *ES) error {
	fmt.Println("AL::EdgeScanEnd")
	return nil
}
func (al *AL) EdgeScanNext(es *ES) error {
	fmt.Println("AL::EdgeScanNext")
	return nil
}


//
// main driver
//
func main() {
	fmt.Println("hello")
	var anyType interface{}
	anyType = 77.0
	fmt.Println("atype:",anyType)
	anyType = "I am a string now"
	fmt.Println("atype:",anyType)

	printAnyType("The car is slow")
	m := map[string] string{"ID":"12345", "name":"Kerry"}
	printAnyType(m)
	printAnyType(12345678)

	// matrix
	gm,_ := NewG()
	gm.MakeG(0)
	gm.AddVertex(1)
	gm.ConnectVertex(1,2,3)
	gm.DisconnectVertex(1,2)
	// list 
	gl,_ := NewG()
	gl.MakeG(1)
	gl.AddVertex(1)
	gl.ConnectVertex(1,2,3)
	gl.DisconnectVertex(1,2)
} // eof main


func printAnyType(val interface{}) {
	fmt.Println("print any value:",val)
}


