package task2

func GetSmallEnvelope(en1 Envelope, en2 Envelope) int{

	if en1.isBigest(en2) {
		return 1
	} else if en2.isBigest(en1){
		return 2
	} else {
		return 0
	}

}
type Envelope struct {
	S1 float64
	S2 float64
}
func (en1 * Envelope) isBigest(en2 Envelope) bool{
	if (en1.S1 > en2.S1 && en1.S2>en2.S2) || (en1.S1>en2.S2 && en1.S2>en2.S1){
		return true
	} else{ return false}
}
