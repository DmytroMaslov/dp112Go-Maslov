//task2
//{"En1":{"S1":4, "S2":5},"En2":{"S1":10, "S2":20}}
package task2
import (
	"errors"
	"encoding/json"
)

type Envelope struct {
	S1 float64 `json:"S1"`
	S2 float64 `json:"S2"`
}
type Params struct {
	En1 Envelope `json:"En1"`
	En2 Envelope `json:"En2"`
}

func Run(param []byte) (string, error){

	var envs = new (Params)
	var err error
	err = json.Unmarshal(param, &envs)
	if err != nil{
		return "", errors.New("can't unmarshal data")
	}
	//err = IsValid(pr[0].En1, pr[1].En2)
	err = IsValid(envs.En1, envs.En2)
	if err != nil{
		return "", err
	}

	if envs.En1.isBigest(envs.En2) {
		return "1", nil
	}
	if envs.En2.isBigest(envs.En1){
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



