package main

import(
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello","world"},
		},
		// add more cases here
	}


	for _, c := range cases{
		actual := cleanInput(c.input)
		lenactual := len(actual)
		lenexpect := len(c.expected)

		if lenactual != lenexpect {
			//t.Errorf("Actual length: %d , does not match expected length: %d", lenactual, lenexpect)
			return 
		}

		for i := range actual {
			word:=actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				//t.Errorf("Words not matching")
				return
			}
		}
	}


	return
}