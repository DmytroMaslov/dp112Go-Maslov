package task4

import "testing"


func TestGetMaxPalidrom(t* testing.T){

 var testData = [] struct{
		input uint64
		expectedResult string
		expectedStatus bool
	}{
	 {123321, "123321", true},
	 {22, "22", true},
	 {123, "", false},
	}

	for _, test := range testData{
		out, status := GetMaxPalidrom(test.input)
		if out != test.expectedResult || status != test.expectedStatus {
			t.Errorf("GetMaxPalidrom(%d) status %s  Result: %s, %s", test.input, test.expectedStatus, out, status)
		}
	}

}
