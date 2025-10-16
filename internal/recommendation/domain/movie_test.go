package domain

import "testing"

func TestFreqId(t *testing.T) {

	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "The man",
			output: "1#0#0#0#1#0#0#1#0#0#0#0#1#1#0#0#0#0#0#1#0#0#0#0#0#0#",
		},
		{
			input:  "Lakhshya",
			output: "2#0#0#0#0#0#0#2#0#0#1#1#0#0#0#0#0#0#1#0#0#0#0#0#1#0#",
		},
	}

	for _, tt := range cases {
		t.Run("Generate FreqId", func(t *testing.T) {
			got := GenerateFreqId(tt.input)

			if got != tt.output {
				t.Errorf("Expected %s , Got %s", tt.output, got)
			}
		})
	}

}
