package ass1

func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

}
