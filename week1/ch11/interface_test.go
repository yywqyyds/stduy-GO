package ch11

import "testing"

type GoProgammer struct{

}

func (g *GoProgammer)WriteHelloWorld() string {
	return "Hello World"
}

type Progammer interface{
	WriteHelloWorld() string
}

func TestClient(t *testing.T){
	var prog Progammer = new(GoProgammer)
	t.Log(prog.WriteHelloWorld())
}