package task2

import "errors"

type Envelope struct {
	S1 float64
	S2 float64
}
func (en1 * Envelope) isBigest(en2 Envelope) bool{
	isParallel := en1.S1 > en2.S1 && en1.S2>en2.S2
	isPerpendicularly := en1.S1>en2.S2 && en1.S2>en2.S1
	if isParallel || isPerpendicularly{
		return true
	}
	return false
}

func IsValid (en1 Envelope, en2 Envelope) (error){
	if (en1.S1 <= 0|| en1.S2<=0){
		return errors.New("Incorrect first envelope")
	}
	if (en2.S1<= 0|| en2.S2<= 0){
		return errors.New("Incorrect second envelope")
	}
	return nil
}

func GetBigestEnvelope(en1 Envelope, en2 Envelope) int{

	if en1.isBigest(en2) {
		return 1
	}
	if en2.isBigest(en1){
		return 2
	}
		return 0

}


