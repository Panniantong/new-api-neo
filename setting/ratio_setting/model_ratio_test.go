package ratio_setting

import "testing"

func TestGetHardcodedCompletionModelRatioGPT56(t *testing.T) {
	tests := []struct {
		name  string
		ratio float64
	}{
		{name: "gpt-5.6", ratio: 6},
		{name: "gpt-5.6-sol", ratio: 6},
		{name: "gpt-5.6-terra", ratio: 6},
		{name: "gpt-5.6-luna", ratio: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ratio, locked := getHardcodedCompletionModelRatio(tt.name)
			if ratio != tt.ratio {
				t.Fatalf("ratio = %v, want %v", ratio, tt.ratio)
			}
			if !locked {
				t.Fatal("GPT-5.6 completion ratio must be locked to official pricing")
			}
		})
	}
}

func TestGetHardcodedCompletionModelRatioExistingGPT5Families(t *testing.T) {
	tests := []struct {
		name  string
		ratio float64
	}{
		{name: "gpt-5.5", ratio: 6},
		{name: "gpt-5.4", ratio: 6},
		{name: "gpt-5.4-nano", ratio: 6.25},
		{name: "gpt-5.3-codex", ratio: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ratio, locked := getHardcodedCompletionModelRatio(tt.name)
			if ratio != tt.ratio {
				t.Fatalf("ratio = %v, want %v", ratio, tt.ratio)
			}
			if !locked {
				t.Fatal("GPT-5 completion ratio must remain locked")
			}
		})
	}
}
