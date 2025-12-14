package main

import "testing"

var day1Input = []Dial{
	{Left, 68},
	{Left, 30},
	{Right, 48},
	{Left, 5},
	{Right, 60},
	{Left, 55},
	{Left, 1},
	{Left, 99},
	{Right, 14},
	{Left, 82},
}

func TestDail(t *testing.T) {
	tests := []struct {
		input         Dial
		current, want int
	}{
		{day1Input[0], 50, 82},
		{day1Input[1], 82, 52},
		{day1Input[2], 52, 0},
		{day1Input[3], 0, 95},
		{day1Input[4], 95, 55},
		{day1Input[5], 55, 0},
		{day1Input[6], 0, 99},
		{day1Input[7], 99, 0},
		{day1Input[8], 0, 14},
		{day1Input[9], 14, 32},
	}

	for _, tt := range tests {
		res := tt.input.Dial(tt.current)
		if res != tt.want {
			t.Errorf("Dial(%d) = %d; want %d", tt.current, res, tt.want)
		}
	}
}

func TestDialPassword(t *testing.T) {
	tests := []struct {
		inputs        []Dial
		current, want int
		part1         bool
	}{
		// {day1Input, 50, 3, true},
		{day1Input, 50, 6, false},
	}

	for _, tt := range tests {
		actual := DialPassword(tt.inputs, tt.current, tt.part1)
		if actual != tt.want {
			t.Errorf("DialPassword(...) = %d; want %d", actual, tt.want)
		}
	}
}

func TestDialZeroCrossings(t *testing.T) {
	tests := []struct {
		uut           Dial
		current, want int
	}{
		{Dial{Left, 2}, 1, 1},
		{Dial{Right, 2}, 99, 1},
		{Dial{Right, 600}, 0, 6},
		{Dial{Left, 450}, 50, 5},
		{day1Input[0], 50, 1},
		{day1Input[1], 82, 0},
		{day1Input[2], 52, 1},
		{day1Input[3], 0, 0},
		{day1Input[4], 95, 1},
		{day1Input[5], 55, 1},
		{day1Input[6], 0, 0},
		{day1Input[7], 99, 1},
		{day1Input[8], 0, 0},
		{day1Input[9], 14, 1},
	}

	for _, tt := range tests {
		actual := tt.uut.DialZeroCrossings(tt.current)
		if actual != tt.want {
			t.Errorf("DialZeroCrossings(Current: %d, Direction: %d, Amount: %d) = Actual: %d; want: %d", tt.current, tt.uut.Direction, tt.uut.Amount, actual, tt.want)
		}
	}
}
