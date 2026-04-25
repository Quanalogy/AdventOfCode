package main

import (
	"reflect"
	"testing"
)

func Test_valid(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Inserting repeating full length should fail", args{"123123"}, false},
		{"Starting with zero should fail", args{"0123"}, false},
		{"Non-repeating should pass", args{"123"}, true},
		{"Odd length should pass", args{"12345"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid(tt.args.x); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInvalidInRange(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Invalid range 11-22 with two invalid numbers should return 11,22", args{"11-22"}, []string{"11", "22"}},
		{"Invalid range 95-115 with one invalid number should return 99", args{"95-115"}, []string{"99"}},
		{"Invalid range 998-1012 with one invalid number should return 1010", args{"998-1012"}, []string{"1010"}},
		{"Invalid range 1188511880-1188511890 with one invalid number should return 1188511885", args{"1188511880-1188511890"}, []string{"1188511885"}},
		{"Invalid range 222220-222224 with one invalid number should return 222222", args{"222220-222224"}, []string{"222222"}},
		{"Valid range 1698522-1698528 should return empty list", args{"1698522-1698528"}, []string{}},
		{"Invalid range 446443-446449 should return 446446", args{"446443-446449"}, []string{"446446"}},
		{"Invalid range 38593856-38593862 should return 38593859", args{"38593856-38593862"}, []string{"38593859"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInvalidInRange(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInvalidInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
