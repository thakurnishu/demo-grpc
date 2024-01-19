package main

const (
	port     = "3000"
	protocol = "tcp"
)

func main() {
	s := NewServer(protocol, port)
	s.Run()
}
