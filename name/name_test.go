package name_test

import (
	"github.com/Jumpaku/sql-gogen-lib/name"
	"testing"
)

func TestName_AllUpperKebab(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "B-BA1C"},
		{in: "c1BbC1", want: "C1-BB-C1"},
		{in: "-a", want: "A"},
		{in: "bBbb1C", want: "B-BBB1-C"},
		{in: "_ a", want: "A"},
		{in: "1", want: "1"},
		{in: "C1", want: "C1"},
		{in: " A_", want: "A"},
		{in: "11C", want: "11-C"},
		{in: "-_bb1", want: "BB1"},
		{in: "a", want: "A"},
		{in: "1c1c1", want: "1C1C1"},
		{in: "bB1a ", want: "B-B1A"},
		{in: "a-1c1", want: "A-1C1"},
		{in: "1C1C11", want: "1-C1-C11"},
		{in: "_a", want: "A"},
		{in: "aC1c1a", want: "A-C1C1A"},
		{in: "bb 1Cc1", want: "BB-1-CC1"},
		{in: "1A-bB", want: "1-A-B-B"},
		{in: "AC1", want: "AC1"},
		{in: "c1", want: "C1"},
		{in: " BBc1", want: "BBC1"},
		{in: "1C", want: "1-C"},
		{in: "c11cc1Bb", want: "C11CC1-BB"},
		{in: "1cbb1cBb", want: "1CBB1C-BB"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "BBC1"},
		{in: "BBa", want: "BBA"},
		{in: "BbbB", want: "BBB-B"},
		{in: "-c1a", want: "C1A"},
		{in: "_ ABB", want: "ABB"},
		{in: "a1C", want: "A1-C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1-CA"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1-CA"},
		{in: " bb1C", want: "BB1-C"},
		{in: "bBBB1a", want: "B-BBB1A"},
		{in: " 1CBb", want: "1-CBB"},
		{in: "bB", want: "B-B"},
		{in: "a--", want: "A"},
		{in: " C1", want: "C1"},
		{in: "1Cc1A1", want: "1-CC1-A1"},
		{in: "bbBB", want: "BB-BB"},
		{in: "BB", want: "BB"},
		{in: "_c1", want: "C1"},
		{in: "1c", want: "1C"},
		{in: "", want: ""},
		{in: "ABB", want: "ABB"},
		{in: "1bb1cc1", want: "1BB1CC1"},
		{in: "BBc1", want: "BBC1"},
		{in: "C11C", want: "C11-C"},
		{in: "BBbbAc1", want: "BBBB-AC1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "BB-BBA1"},
		{in: "A1C", want: "A1-C"},
		{in: "A", want: "A"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "C1"},
		{in: "Bbc1Bb", want: "BBC1-BB"},
		{in: "bb", want: "BB"},
		{in: "bBBb-", want: "B-BBB"},
		{in: "Bbc1bb", want: "BBC1BB"},
		{in: "-Bb", want: "BB"},
		{in: " bB", want: "B-B"},
		{in: "1-1", want: "1-1"},
		{in: "-c1-", want: "C1"},
		{in: " 1bBa", want: "1B-BA"},
		{in: "Bb", want: "BB"},
		{in: "1CbBBB", want: "1-CB-BBB"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			sut := name.New(tt.in)
			if got := sut.AllUpperKebab(); got != tt.want {
				t.Errorf("AllUpperKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_FirstUpperKebab(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "B-Ba1c"},
		{in: "c1BbC1", want: "C1-Bb-C1"},
		{in: "-a", want: "A"},
		{in: "bBbb1C", want: "B-Bbb1-C"},
		{in: "_ a", want: "A"},
		{in: "1", want: "1"},
		{in: "C1", want: "C1"},
		{in: " A_", want: "A"},
		{in: "11C", want: "11-C"},
		{in: "-_bb1", want: "Bb1"},
		{in: "a", want: "A"},
		{in: "1c1c1", want: "1c1c1"},
		{in: "bB1a ", want: "B-B1a"},
		{in: "a-1c1", want: "A-1c1"},
		{in: "1C1C11", want: "1-C1-C11"},
		{in: "_a", want: "A"},
		{in: "aC1c1a", want: "A-C1c1a"},
		{in: "bb 1Cc1", want: "Bb-1-Cc1"},
		{in: "1A-bB", want: "1-A-B-B"},
		{in: "AC1", want: "Ac1"},
		{in: "c1", want: "C1"},
		{in: " BBc1", want: "Bbc1"},
		{in: "1C", want: "1-C"},
		{in: "c11cc1Bb", want: "C11cc1-Bb"},
		{in: "1cbb1cBb", want: "1cbb1c-Bb"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "Bbc1"},
		{in: "BBa", want: "Bba"},
		{in: "BbbB", want: "Bbb-B"},
		{in: "-c1a", want: "C1a"},
		{in: "_ ABB", want: "Abb"},
		{in: "a1C", want: "A1-C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1-Ca"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1-Ca"},
		{in: " bb1C", want: "Bb1-C"},
		{in: "bBBB1a", want: "B-Bbb1a"},
		{in: " 1CBb", want: "1-Cbb"},
		{in: "bB", want: "B-B"},
		{in: "a--", want: "A"},
		{in: " C1", want: "C1"},
		{in: "1Cc1A1", want: "1-Cc1-A1"},
		{in: "bbBB", want: "Bb-Bb"},
		{in: "BB", want: "Bb"},
		{in: "_c1", want: "C1"},
		{in: "1c", want: "1c"},
		{in: "", want: ""},
		{in: "ABB", want: "Abb"},
		{in: "1bb1cc1", want: "1bb1cc1"},
		{in: "BBc1", want: "Bbc1"},
		{in: "C11C", want: "C11-C"},
		{in: "BBbbAc1", want: "Bbbb-Ac1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "Bb-Bba1"},
		{in: "A1C", want: "A1-C"},
		{in: "A", want: "A"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "C1"},
		{in: "Bbc1Bb", want: "Bbc1-Bb"},
		{in: "bb", want: "Bb"},
		{in: "bBBb-", want: "B-Bbb"},
		{in: "Bbc1bb", want: "Bbc1bb"},
		{in: "-Bb", want: "Bb"},
		{in: " bB", want: "B-B"},
		{in: "1-1", want: "1-1"},
		{in: "-c1-", want: "C1"},
		{in: " 1bBa", want: "1b-Ba"},
		{in: "Bb", want: "Bb"},
		{in: "1CbBBB", want: "1-Cb-Bbb"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.FirstUpperKebab(); got != tt.want {
				t.Errorf("FirstUpperKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_LowerKebab(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "b-ba1c"},
		{in: "c1BbC1", want: "c1-bb-c1"},
		{in: "-a", want: "a"},
		{in: "bBbb1C", want: "b-bbb1-c"},
		{in: "_ a", want: "a"},
		{in: "1", want: "1"},
		{in: "C1", want: "c1"},
		{in: " A_", want: "a"},
		{in: "11C", want: "11-c"},
		{in: "-_bb1", want: "bb1"},
		{in: "a", want: "a"},
		{in: "1c1c1", want: "1c1c1"},
		{in: "bB1a ", want: "b-b1a"},
		{in: "a-1c1", want: "a-1c1"},
		{in: "1C1C11", want: "1-c1-c11"},
		{in: "_a", want: "a"},
		{in: "aC1c1a", want: "a-c1c1a"},
		{in: "bb 1Cc1", want: "bb-1-cc1"},
		{in: "1A-bB", want: "1-a-b-b"},
		{in: "AC1", want: "ac1"},
		{in: "c1", want: "c1"},
		{in: " BBc1", want: "bbc1"},
		{in: "1C", want: "1-c"},
		{in: "c11cc1Bb", want: "c11cc1-bb"},
		{in: "1cbb1cBb", want: "1cbb1c-bb"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "bbc1"},
		{in: "BBa", want: "bba"},
		{in: "BbbB", want: "bbb-b"},
		{in: "-c1a", want: "c1a"},
		{in: "_ ABB", want: "abb"},
		{in: "a1C", want: "a1-c"},
		{in: "_", want: ""},
		{in: "1CA", want: "1-ca"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1-ca"},
		{in: " bb1C", want: "bb1-c"},
		{in: "bBBB1a", want: "b-bbb1a"},
		{in: " 1CBb", want: "1-cbb"},
		{in: "bB", want: "b-b"},
		{in: "a--", want: "a"},
		{in: " C1", want: "c1"},
		{in: "1Cc1A1", want: "1-cc1-a1"},
		{in: "bbBB", want: "bb-bb"},
		{in: "BB", want: "bb"},
		{in: "_c1", want: "c1"},
		{in: "1c", want: "1c"},
		{in: "", want: ""},
		{in: "ABB", want: "abb"},
		{in: "1bb1cc1", want: "1bb1cc1"},
		{in: "BBc1", want: "bbc1"},
		{in: "C11C", want: "c11-c"},
		{in: "BBbbAc1", want: "bbbb-ac1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "bb-bba1"},
		{in: "A1C", want: "a1-c"},
		{in: "A", want: "a"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "c1"},
		{in: "Bbc1Bb", want: "bbc1-bb"},
		{in: "bb", want: "bb"},
		{in: "bBBb-", want: "b-bbb"},
		{in: "Bbc1bb", want: "bbc1bb"},
		{in: "-Bb", want: "bb"},
		{in: " bB", want: "b-b"},
		{in: "1-1", want: "1-1"},
		{in: "-c1-", want: "c1"},
		{in: " 1bBa", want: "1b-ba"},
		{in: "Bb", want: "bb"},
		{in: "1CbBBB", want: "1-cb-bbb"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.LowerKebab(); got != tt.want {
				t.Errorf("LowerKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_AllUpperSnake(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "B_BA1C"},
		{in: "c1BbC1", want: "C1_BB_C1"},
		{in: "-a", want: "A"},
		{in: "bBbb1C", want: "B_BBB1_C"},
		{in: "_ a", want: "A"},
		{in: "1", want: "1"},
		{in: "C1", want: "C1"},
		{in: " A_", want: "A"},
		{in: "11C", want: "11_C"},
		{in: "-_bb1", want: "BB1"},
		{in: "a", want: "A"},
		{in: "1c1c1", want: "1C1C1"},
		{in: "bB1a ", want: "B_B1A"},
		{in: "a-1c1", want: "A_1C1"},
		{in: "1C1C11", want: "1_C1_C11"},
		{in: "_a", want: "A"},
		{in: "aC1c1a", want: "A_C1C1A"},
		{in: "bb 1Cc1", want: "BB_1_CC1"},
		{in: "1A-bB", want: "1_A_B_B"},
		{in: "AC1", want: "AC1"},
		{in: "c1", want: "C1"},
		{in: " BBc1", want: "BBC1"},
		{in: "1C", want: "1_C"},
		{in: "c11cc1Bb", want: "C11CC1_BB"},
		{in: "1cbb1cBb", want: "1CBB1C_BB"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "BBC1"},
		{in: "BBa", want: "BBA"},
		{in: "BbbB", want: "BBB_B"},
		{in: "-c1a", want: "C1A"},
		{in: "_ ABB", want: "ABB"},
		{in: "a1C", want: "A1_C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1_CA"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1_CA"},
		{in: " bb1C", want: "BB1_C"},
		{in: "bBBB1a", want: "B_BBB1A"},
		{in: " 1CBb", want: "1_CBB"},
		{in: "bB", want: "B_B"},
		{in: "a--", want: "A"},
		{in: " C1", want: "C1"},
		{in: "1Cc1A1", want: "1_CC1_A1"},
		{in: "bbBB", want: "BB_BB"},
		{in: "BB", want: "BB"},
		{in: "_c1", want: "C1"},
		{in: "1c", want: "1C"},
		{in: "", want: ""},
		{in: "ABB", want: "ABB"},
		{in: "1bb1cc1", want: "1BB1CC1"},
		{in: "BBc1", want: "BBC1"},
		{in: "C11C", want: "C11_C"},
		{in: "BBbbAc1", want: "BBBB_AC1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "BB_BBA1"},
		{in: "A1C", want: "A1_C"},
		{in: "A", want: "A"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "C1"},
		{in: "Bbc1Bb", want: "BBC1_BB"},
		{in: "bb", want: "BB"},
		{in: "bBBb-", want: "B_BBB"},
		{in: "Bbc1bb", want: "BBC1BB"},
		{in: "-Bb", want: "BB"},
		{in: " bB", want: "B_B"},
		{in: "1-1", want: "1_1"},
		{in: "-c1-", want: "C1"},
		{in: " 1bBa", want: "1B_BA"},
		{in: "Bb", want: "BB"},
		{in: "1CbBBB", want: "1_CB_BBB"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.AllUpperSnake(); got != tt.want {
				t.Errorf("AllUpperSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_FirstUpperSnake(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "B_BA1C"},
		{in: "c1BbC1", want: "C1_BB_C1"},
		{in: "-a", want: "A"},
		{in: "bBbb1C", want: "B_BBB1_C"},
		{in: "_ a", want: "A"},
		{in: "1", want: "1"},
		{in: "C1", want: "C1"},
		{in: " A_", want: "A"},
		{in: "11C", want: "11_C"},
		{in: "-_bb1", want: "BB1"},
		{in: "a", want: "A"},
		{in: "1c1c1", want: "1C1C1"},
		{in: "bB1a ", want: "B_B1A"},
		{in: "a-1c1", want: "A_1C1"},
		{in: "1C1C11", want: "1_C1_C11"},
		{in: "_a", want: "A"},
		{in: "aC1c1a", want: "A_C1C1A"},
		{in: "bb 1Cc1", want: "BB_1_CC1"},
		{in: "1A-bB", want: "1_A_B_B"},
		{in: "AC1", want: "AC1"},
		{in: "c1", want: "C1"},
		{in: " BBc1", want: "BBC1"},
		{in: "1C", want: "1_C"},
		{in: "c11cc1Bb", want: "C11CC1_BB"},
		{in: "1cbb1cBb", want: "1CBB1C_BB"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "BBC1"},
		{in: "BBa", want: "BBA"},
		{in: "BbbB", want: "BBB_B"},
		{in: "-c1a", want: "C1A"},
		{in: "_ ABB", want: "ABB"},
		{in: "a1C", want: "A1_C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1_CA"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1_CA"},
		{in: " bb1C", want: "BB1_C"},
		{in: "bBBB1a", want: "B_BBB1A"},
		{in: " 1CBb", want: "1_CBB"},
		{in: "bB", want: "B_B"},
		{in: "a--", want: "A"},
		{in: " C1", want: "C1"},
		{in: "1Cc1A1", want: "1_CC1_A1"},
		{in: "bbBB", want: "BB_BB"},
		{in: "BB", want: "BB"},
		{in: "_c1", want: "C1"},
		{in: "1c", want: "1C"},
		{in: "", want: ""},
		{in: "ABB", want: "ABB"},
		{in: "1bb1cc1", want: "1BB1CC1"},
		{in: "BBc1", want: "BBC1"},
		{in: "C11C", want: "C11_C"},
		{in: "BBbbAc1", want: "BBBB_AC1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "BB_BBA1"},
		{in: "A1C", want: "A1_C"},
		{in: "A", want: "A"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "C1"},
		{in: "Bbc1Bb", want: "BBC1_BB"},
		{in: "bb", want: "BB"},
		{in: "bBBb-", want: "B_BBB"},
		{in: "Bbc1bb", want: "BBC1BB"},
		{in: "-Bb", want: "BB"},
		{in: " bB", want: "B_B"},
		{in: "1-1", want: "1_1"},
		{in: "-c1-", want: "C1"},
		{in: " 1bBa", want: "1B_BA"},
		{in: "Bb", want: "BB"},
		{in: "1CbBBB", want: "1_CB_BBB"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.FirstUpperSnake(); got != tt.want {
				t.Errorf("FirstUpperSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_LowerSnake(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "b_ba1c"},
		{in: "c1BbC1", want: "c1_bb_c1"},
		{in: "-a", want: "a"},
		{in: "bBbb1C", want: "b_bbb1_c"},
		{in: "_ a", want: "a"},
		{in: "1", want: "1"},
		{in: "C1", want: "c1"},
		{in: " A_", want: "a"},
		{in: "11C", want: "11_c"},
		{in: "-_bb1", want: "bb1"},
		{in: "a", want: "a"},
		{in: "1c1c1", want: "1c1c1"},
		{in: "bB1a ", want: "b_b1a"},
		{in: "a-1c1", want: "a_1c1"},
		{in: "1C1C11", want: "1_c1_c11"},
		{in: "_a", want: "a"},
		{in: "aC1c1a", want: "a_c1c1a"},
		{in: "bb 1Cc1", want: "bb_1_cc1"},
		{in: "1A-bB", want: "1_a_b_b"},
		{in: "AC1", want: "ac1"},
		{in: "c1", want: "c1"},
		{in: " BBc1", want: "bbc1"},
		{in: "1C", want: "1_c"},
		{in: "c11cc1Bb", want: "c11cc1_bb"},
		{in: "1cbb1cBb", want: "1cbb1c_bb"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "bbc1"},
		{in: "BBa", want: "bba"},
		{in: "BbbB", want: "bbb_b"},
		{in: "-c1a", want: "c1a"},
		{in: "_ ABB", want: "abb"},
		{in: "a1C", want: "a1_c"},
		{in: "_", want: ""},
		{in: "1CA", want: "1_ca"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1_ca"},
		{in: " bb1C", want: "bb1_c"},
		{in: "bBBB1a", want: "b_bbb1a"},
		{in: " 1CBb", want: "1_cbb"},
		{in: "bB", want: "b_b"},
		{in: "a--", want: "a"},
		{in: " C1", want: "c1"},
		{in: "1Cc1A1", want: "1_cc1_a1"},
		{in: "bbBB", want: "bb_bb"},
		{in: "BB", want: "bb"},
		{in: "_c1", want: "c1"},
		{in: "1c", want: "1c"},
		{in: "", want: ""},
		{in: "ABB", want: "abb"},
		{in: "1bb1cc1", want: "1bb1cc1"},
		{in: "BBc1", want: "bbc1"},
		{in: "C11C", want: "c11_c"},
		{in: "BBbbAc1", want: "bbbb_ac1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "bb_bba1"},
		{in: "A1C", want: "a1_c"},
		{in: "A", want: "a"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "c1"},
		{in: "Bbc1Bb", want: "bbc1_bb"},
		{in: "bb", want: "bb"},
		{in: "bBBb-", want: "b_bbb"},
		{in: "Bbc1bb", want: "bbc1bb"},
		{in: "-Bb", want: "bb"},
		{in: " bB", want: "b_b"},
		{in: "1-1", want: "1_1"},
		{in: "-c1-", want: "c1"},
		{in: " 1bBa", want: "1b_ba"},
		{in: "Bb", want: "bb"},
		{in: "1CbBBB", want: "1_cb_bbb"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.LowerSnake(); got != tt.want {
				t.Errorf("LowerSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_UpperCamel(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "BBa1c"},
		{in: "c1BbC1", want: "C1BbC1"},
		{in: "-a", want: "A"},
		{in: "bBbb1C", want: "BBbb1C"},
		{in: "_ a", want: "A"},
		{in: "1", want: "1"},
		{in: "C1", want: "C1"},
		{in: " A_", want: "A"},
		{in: "11C", want: "11C"},
		{in: "-_bb1", want: "Bb1"},
		{in: "a", want: "A"},
		{in: "1c1c1", want: "1c1c1"},
		{in: "bB1a ", want: "BB1a"},
		{in: "a-1c1", want: "A1c1"},
		{in: "1C1C11", want: "1C1C11"},
		{in: "_a", want: "A"},
		{in: "aC1c1a", want: "AC1c1a"},
		{in: "bb 1Cc1", want: "Bb1Cc1"},
		{in: "1A-bB", want: "1ABB"},
		{in: "AC1", want: "Ac1"},
		{in: "c1", want: "C1"},
		{in: " BBc1", want: "Bbc1"},
		{in: "1C", want: "1C"},
		{in: "c11cc1Bb", want: "C11cc1Bb"},
		{in: "1cbb1cBb", want: "1cbb1cBb"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "Bbc1"},
		{in: "BBa", want: "Bba"},
		{in: "BbbB", want: "BbbB"},
		{in: "-c1a", want: "C1a"},
		{in: "_ ABB", want: "Abb"},
		{in: "a1C", want: "A1C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1Ca"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1Ca"},
		{in: " bb1C", want: "Bb1C"},
		{in: "bBBB1a", want: "BBbb1a"},
		{in: " 1CBb", want: "1Cbb"},
		{in: "bB", want: "BB"},
		{in: "a--", want: "A"},
		{in: " C1", want: "C1"},
		{in: "1Cc1A1", want: "1Cc1A1"},
		{in: "bbBB", want: "BbBb"},
		{in: "BB", want: "Bb"},
		{in: "_c1", want: "C1"},
		{in: "1c", want: "1c"},
		{in: "", want: ""},
		{in: "ABB", want: "Abb"},
		{in: "1bb1cc1", want: "1bb1cc1"},
		{in: "BBc1", want: "Bbc1"},
		{in: "C11C", want: "C11C"},
		{in: "BBbbAc1", want: "BbbbAc1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "BbBba1"},
		{in: "A1C", want: "A1C"},
		{in: "A", want: "A"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "C1"},
		{in: "Bbc1Bb", want: "Bbc1Bb"},
		{in: "bb", want: "Bb"},
		{in: "bBBb-", want: "BBbb"},
		{in: "Bbc1bb", want: "Bbc1bb"},
		{in: "-Bb", want: "Bb"},
		{in: " bB", want: "BB"},
		{in: "1-1", want: "11"},
		{in: "-c1-", want: "C1"},
		{in: " 1bBa", want: "1bBa"},
		{in: "Bb", want: "Bb"},
		{in: "1CbBBB", want: "1CbBbb"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.UpperCamel(); got != tt.want {
				t.Errorf("UpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_LowerCamel(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "bBA1c_", want: "bBa1c"},
		{in: "c1BbC1", want: "c1BbC1"},
		{in: "-a", want: "a"},
		{in: "bBbb1C", want: "bBbb1C"},
		{in: "_ a", want: "a"},
		{in: "1", want: "1"},
		{in: "C1", want: "c1"},
		{in: " A_", want: "a"},
		{in: "11C", want: "11C"},
		{in: "-_bb1", want: "bb1"},
		{in: "a", want: "a"},
		{in: "1c1c1", want: "1c1c1"},
		{in: "bB1a ", want: "bB1a"},
		{in: "a-1c1", want: "a1c1"},
		{in: "1C1C11", want: "1C1C11"},
		{in: "_a", want: "a"},
		{in: "aC1c1a", want: "aC1c1a"},
		{in: "bb 1Cc1", want: "bb1Cc1"},
		{in: "1A-bB", want: "1ABB"},
		{in: "AC1", want: "ac1"},
		{in: "c1", want: "c1"},
		{in: " BBc1", want: "bbc1"},
		{in: "1C", want: "1C"},
		{in: "c11cc1Bb", want: "c11cc1Bb"},
		{in: "1cbb1cBb", want: "1cbb1cBb"},
		{in: " ", want: ""},
		{in: "__BBC1", want: "bbc1"},
		{in: "BBa", want: "bba"},
		{in: "BbbB", want: "bbbB"},
		{in: "-c1a", want: "c1a"},
		{in: "_ ABB", want: "abb"},
		{in: "a1C", want: "a1C"},
		{in: "_", want: ""},
		{in: "1CA", want: "1Ca"},
		{in: "--_", want: ""},
		{in: "1Ca", want: "1Ca"},
		{in: " bb1C", want: "bb1C"},
		{in: "bBBB1a", want: "bBbb1a"},
		{in: " 1CBb", want: "1Cbb"},
		{in: "bB", want: "bB"},
		{in: "a--", want: "a"},
		{in: " C1", want: "c1"},
		{in: "1Cc1A1", want: "1Cc1A1"},
		{in: "bbBB", want: "bbBb"},
		{in: "BB", want: "bb"},
		{in: "_c1", want: "c1"},
		{in: "1c", want: "1c"},
		{in: "", want: ""},
		{in: "ABB", want: "abb"},
		{in: "1bb1cc1", want: "1bb1cc1"},
		{in: "BBc1", want: "bbc1"},
		{in: "C11C", want: "c11C"},
		{in: "BBbbAc1", want: "bbbbAc1"},
		{in: " 1", want: "1"},
		{in: "bbBBA1", want: "bbBba1"},
		{in: "A1C", want: "a1C"},
		{in: "A", want: "a"},
		{in: "-_", want: ""},
		{in: "_C1_", want: "c1"},
		{in: "Bbc1Bb", want: "bbc1Bb"},
		{in: "bb", want: "bb"},
		{in: "bBBb-", want: "bBbb"},
		{in: "Bbc1bb", want: "bbc1bb"},
		{in: "-Bb", want: "bb"},
		{in: " bB", want: "bB"},
		{in: "1-1", want: "11"},
		{in: "-c1-", want: "c1"},
		{in: " 1bBa", want: "1bBa"},
		{in: "Bb", want: "bb"},
		{in: "1CbBBB", want: "1CbBbb"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			n := name.New(tt.in)
			if got := n.LowerCamel(); got != tt.want {
				t.Errorf("LowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
