package main

import (
	"fmt"

	t "github.com/sap200/nfa/type"
)

func main() {
	Q := []string{"1", "2", "3", "4"}
	A := []string{"0", "1"}
	F := []string{"4"}
	S := "1"
	trans_fun := [][3]string{
		[3]string{"1", "0", "1"},
		[3]string{"1", "0", "2"},
		[3]string{"1", "1", "1"},
		[3]string{"1", "1", "3"},
		[3]string{"2", "0", "4"},
		[3]string{"3", "1", "4"},
		[3]string{"4", "0", "4"},
		[3]string{"4", "1", "4"},
	}

	nfa := t.New_NFA(Q, A, S, F, trans_fun)

	fmt.Println("\nIs string accepted? ", nfa.Is_Seq_Accepted("11001101", false))

	seqs := t.Generate_All(A, 10)
	for _, seq := range seqs {
		if nfa.Is_Seq_Accepted(seq, false) {
			fmt.Println(seq)
		}
	}

}
