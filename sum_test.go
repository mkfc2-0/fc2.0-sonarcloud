package main

import "testing"

func TestSum(t *testing.T) {
	soma := sum(2, 3)

	if soma != 5 {
		t.Error("Soma diferente de 5.")
	}
}

func TestSub(t *testing.T) {

	subtracao := sub(4, 3)

	if subtracao != 1 {
		t.Error("Subtração diferente de 1.")
	}
}

func TestTimes(t *testing.T) {

	multiplica := times(4, 3)

	if multiplica != 12 {
		t.Error("Multiplicação diferente de 12.")
	}
}

func TestSumX(t *testing.T) {

	somaX := sumX(4, 3)

	if somaX != 11 {
		t.Error("SomaX diferente de 11.")
	}
}
