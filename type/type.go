package nfa

import (
	"fmt"
)

type NFA struct {
	state_set_Q       []string
	input_set_A       []string
	initial_state_q0  string
	final_state_set_F []string
	transition_func   [][3]string
}

func New_NFA(state_set_Q []string, input_set_A []string, initial_state_q0 string, final_state_set_F []string, transition_func [][3]string) NFA {
	nfa := NFA{
		state_set_Q:       state_set_Q,
		input_set_A:       input_set_A,
		initial_state_q0:  initial_state_q0,
		final_state_set_F: final_state_set_F,
		transition_func:   transition_func,
	}

	return nfa
}

func (nfa NFA) Next_State(current_state string, current_input string) []string {
	trans_fun := nfa.transition_func
	var next_states []string

	for _, value := range trans_fun {
		init_state := value[0]
		input_alphabet := value[1]
		out_state := value[2]
		if init_state == current_state && current_input == input_alphabet {
			next_states = append(next_states, out_state)
		}
	}

	return next_states
}

func (nfa NFA) Is_Seq_Accepted(input_seq string, verbose bool) bool {
	current_state := []string{nfa.initial_state_q0}

	if verbose {
		fmt.Println("inp_seq\t\t  state|inp\trem_seq")
		fmt.Println("-------------------------------------------")
	}

	for i := 0; i < len(input_seq); i++ {
		current_input := string(input_seq[i])
		if verbose {
			fmt.Println(input_seq[i:], "\t\t{ ", current_state, "|", current_input, "}\t", input_seq[i+1:])
		}
		var pos_next_states []string
		for j := 0; j < len(current_state); j++ {
			pos_next_states = append(pos_next_states, nfa.Next_State(current_state[j], current_input)...)
		}
		current_state = pos_next_states
	}

	for _, el := range current_state {
		for _, fs := range nfa.final_state_set_F {
			if el == fs {
				return true
			}
		}
	}

	return false
}

func Generate_All(alphabet_set []string, size int) []string {
	if size == 0 {
		return []string{}
	} else if size == 1 {
		return alphabet_set
	} else {
		return Cartesian_Product(alphabet_set, Generate_All(alphabet_set, size-1))
	}
}

func Cartesian_Product(set1 []string, set2 []string) []string {
	var result_set []string
	for _, val1 := range set1 {
		for _, val2 := range set2 {
			result_set = append(result_set, val1+val2)
		}
	}

	return result_set
}
