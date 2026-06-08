package sevaluation

import (
	"testing"

	"github.com/plexusone/omniobserve/llmops"
	"github.com/plexusone/structured-evaluation/rubric"
)

func TestNormalizeScore(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0.0, 0.0},
		{5.0, 0.5},
		{7.0, 0.7},
		{10.0, 1.0},
	}

	for _, tt := range tests {
		got := NormalizeScore(tt.input)
		if got != tt.expected {
			t.Errorf("NormalizeScore(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestDenormalizeScore(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0.0, 0.0},
		{0.5, 5.0},
		{0.7, 7.0},
		{1.0, 10.0},
	}

	for _, tt := range tests {
		got := DenormalizeScore(tt.input)
		if got != tt.expected {
			t.Errorf("DenormalizeScore(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestFormatFinding(t *testing.T) {
	f := rubric.Finding{
		Severity:       rubric.SeverityHigh,
		Category:       "security",
		Title:          "SQL Injection Risk",
		Recommendation: "Use parameterized queries",
	}

	got := formatFinding(f)
	expected := "[high] security: SQL Injection Risk - Use parameterized queries"

	if got != expected {
		t.Errorf("formatFinding() = %v, want %v", got, expected)
	}
}

func TestDefaultExportOptions(t *testing.T) {
	opts := DefaultExportOptions()

	if !opts.IncludeFindings {
		t.Error("IncludeFindings should default to true")
	}
	if !opts.IncludeCategories {
		t.Error("IncludeCategories should default to true")
	}
	if !opts.IncludeOverall {
		t.Error("IncludeOverall should default to true")
	}
	if opts.ScorePrefix != "" {
		t.Error("ScorePrefix should default to empty string")
	}
	if opts.Source != "structured-evaluation" {
		t.Errorf("Source = %v, want structured-evaluation", opts.Source)
	}
}

func TestImportEvalResult(t *testing.T) {
	result := &llmops.EvalResult{
		Scores: []llmops.MetricScore{
			{Name: "relevance", Score: 0.8, Reason: "Good relevance"},
			{Name: "coherence", Score: 0.6, Reason: "Some issues"},
		},
	}

	report := ImportEvalResult(result)

	if len(report.Categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(report.Categories))
	}

	if report.Categories[0].Category != "relevance" {
		t.Errorf("Expected category 'relevance', got %s", report.Categories[0].Category)
	}

	// 0.8 normalized -> 8.0 denormalized, stored in NumericScore
	if report.Categories[0].NumericScore == nil {
		t.Fatal("Expected NumericScore to be set")
	}
	if *report.Categories[0].NumericScore != 8.0 {
		t.Errorf("Expected NumericScore 8.0, got %f", *report.Categories[0].NumericScore)
	}

	// Categorical score should be pass (>= 7.0)
	if report.Categories[0].Score != rubric.ScorePass {
		t.Errorf("Expected Score pass, got %s", report.Categories[0].Score)
	}
}

func TestImportEvalResultWithError(t *testing.T) {
	result := &llmops.EvalResult{
		Scores: []llmops.MetricScore{
			{Name: "test", Score: 0.5, Error: "evaluation failed"},
		},
	}

	report := ImportEvalResult(result)

	if len(report.Findings) != 1 {
		t.Errorf("Expected 1 finding for error, got %d", len(report.Findings))
	}

	if report.Findings[0].Title != "Evaluation error" {
		t.Errorf("Expected finding title 'Evaluation error', got %s", report.Findings[0].Title)
	}
}

func TestMetricScoreToCategory(t *testing.T) {
	score := llmops.MetricScore{
		Name:   "accuracy",
		Score:  0.75,
		Reason: "Good accuracy",
	}

	cat := MetricScoreToCategory(score)

	if cat.Category != "accuracy" {
		t.Errorf("Expected category 'accuracy', got %s", cat.Category)
	}
	// 0.75 normalized -> 7.5 denormalized
	if cat.NumericScore == nil {
		t.Fatal("Expected NumericScore to be set")
	}
	if *cat.NumericScore != 7.5 {
		t.Errorf("Expected NumericScore 7.5, got %f", *cat.NumericScore)
	}
	// 7.5 >= 7.0 -> pass
	if cat.Score != rubric.ScorePass {
		t.Errorf("Expected Score pass, got %s", cat.Score)
	}
}

func TestAnnotationToFinding(t *testing.T) {
	ann := llmops.Annotation{
		Name:        "Security Issue",
		Label:       "high",
		Explanation: "Found vulnerability",
		Metadata: map[string]any{
			"category": "security",
		},
	}

	f := AnnotationToFinding(ann)

	if f.Severity != rubric.SeverityHigh {
		t.Errorf("Expected severity high, got %s", f.Severity)
	}
	if f.Category != "security" {
		t.Errorf("Expected category 'security', got %s", f.Category)
	}
	if f.Title != "Security Issue" {
		t.Errorf("Expected title 'Security Issue', got %s", f.Title)
	}
}

func TestAnnotationToFinding_SeverityMapping(t *testing.T) {
	tests := []struct {
		label    string
		expected rubric.Severity
	}{
		{"critical", rubric.SeverityCritical},
		{"CRITICAL", rubric.SeverityCritical},
		{"high", rubric.SeverityHigh},
		{"HIGH", rubric.SeverityHigh},
		{"medium", rubric.SeverityMedium},
		{"MEDIUM", rubric.SeverityMedium},
		{"low", rubric.SeverityLow},
		{"LOW", rubric.SeverityLow},
		{"unknown", rubric.SeverityInfo},
		{"", rubric.SeverityInfo},
	}

	for _, tt := range tests {
		ann := llmops.Annotation{Label: tt.label}
		f := AnnotationToFinding(ann)
		if f.Severity != tt.expected {
			t.Errorf("Label %q: got severity %s, want %s", tt.label, f.Severity, tt.expected)
		}
	}
}
