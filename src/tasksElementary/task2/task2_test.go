package task2

import (
	"testing"
)

func TestGetSmallEnvelope(t *testing.T) {
	var envelopTestData = []struct {
		en1 Envelope
		en2 Envelope
		expectedResult int
	}{
		{Envelope{S1: 5, S2: 10}, Envelope{S1:10, S2:20}, 2},
		{Envelope{S1: 50, S2: 100}, Envelope{S1:1, S2:2}, 1},
		{Envelope{S1: 5, S2: 10}, Envelope{S1:5, S2:10}, 0},
		{Envelope{S1: 50, S2: 5}, Envelope{S1:4, S2:49}, 1},
		{Envelope{S1: 5, S2: 10}, Envelope{S1:10, S2:5.1}, 0},
		{Envelope{S1: 0.00015, S2: 0.00020}, Envelope{S1:0.001, S2:0.002}, 2},
	}
	for _,en := range envelopTestData{
		smallEnvelop := GetBigestEnvelope(en.en1, en.en2)
		if smallEnvelop != en.expectedResult{
			t.Errorf("Input: envelop1: %v envelop2: %v expected: %d actual: %d",en.en1, en.en2,en.expectedResult, smallEnvelop)
		}
	}
}

func TestIsValid(t *testing.T) {
	var envelopTestData = []struct {
		en1 Envelope
		en2 Envelope
		errorMessage string
	}{
		{Envelope{S1: -5, S2: 10}, Envelope{S1:10, S2:20}, "Incorrect first envelope"},
		{Envelope{S1: 0, S2: 10}, Envelope{S1:10, S2:20}, "Incorrect first envelope"},
		{Envelope{S1: 1, S2: 10}, Envelope{S1:-10, S2:20}, "Incorrect second envelope"},
	}
	for _, en := range envelopTestData{
		er := IsValid(en.en1, en.en2)
		if er == nil {
			t.Errorf("Input: envelop1: %v envelop2: %v expected: Error actual: nil",en.en1, en.en2)
		}
		if er !=nil && er.Error() != en.errorMessage{
			t.Errorf("Input: envelop1: %v envelop2: %v expected:%s actual:%s ",en.en1, en.en2, en.errorMessage, er.Error())
		}
	}
}