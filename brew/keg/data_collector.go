package keg

type Writer interface {
	write(line string)
}

func CreateCollector(control KegControl, writer Writer) {

}
