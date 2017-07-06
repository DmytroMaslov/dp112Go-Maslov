//task2
//{"En1":{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
package task2
import (
	"errors"
	"encoding/json"
)

type Envelope struct {
	S1, S2 float64
}
type Params struct {
	En1, En2 Envelope
}

func Run(param []byte) (string, error){
	var pr  = new (Params)
	err := json.Unmarshal(param, &pr)
	if err != nil{
		return "", errors.New("can't unmarshal data")
	}
	err = IsValid(pr.En1, pr.En2)
	if err != nil{
		return "", err
	}

	if pr.En1.isBigest(pr.En2) {
		return "1", nil
	}
	if pr.En2.isBigest(pr.En1){
		return "2", nil
	}
	return "0", nil

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



