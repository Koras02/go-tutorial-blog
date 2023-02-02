func Fprintln(w.io.writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a, true, true)
	n, err = w.Write(p.buff)
	p.free()
	return
}