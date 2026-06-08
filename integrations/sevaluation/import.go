package sevaluation

import (
	"github.com/plexusone/omniobserve/llmops"
	"github.com/plexusone/structured-evaluation/rubric"
)

// ImportOptions configures the import behavior.
type ImportOptions struct {
	// ReviewType sets the review type for the generated report.
	// Default: "llm_evaluation"
	ReviewType string

	// Document sets the document name for the report metadata.
	// Default: ""
	Document string

	// PassCriteria sets the criteria for the report.
	// Default: rubric.DefaultPassCriteria()
	PassCriteria rubric.PassCriteria
}

// DefaultImportOptions returns the default import configuration.
func DefaultImportOptions() ImportOptions {
	return ImportOptions{
		ReviewType:   "llm_evaluation",
		Document:     "",
		PassCriteria: rubric.DefaultPassCriteria(),
	}
}

// ImportEvalResult converts an llmops EvalResult into a Rubric report.
// Each MetricScore becomes a CategoryResult in the report.
func ImportEvalResult(result *llmops.EvalResult, opts ...ImportOptions) *rubric.Rubric {
	opt := DefaultImportOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}

	report := rubric.NewRubric(opt.ReviewType, opt.Document)
	report.PassCriteria = opt.PassCriteria

	for _, score := range result.Scores {
		numericScore := denormalizeScore(score.Score)
		cat := rubric.CategoryResult{
			Category:     score.Name,
			Score:        numericScoreToScoreValue(numericScore),
			NumericScore: &numericScore,
			Reasoning:    score.Reason,
		}
		report.AddCategoryResult(cat)

		// If score has an error, add it as a finding
		if score.Error != "" {
			report.AddFinding(rubric.Finding{
				Severity:       rubric.SeverityMedium,
				Category:       score.Name,
				Title:          "Evaluation error",
				Description:    score.Error,
				Recommendation: "Review and re-run evaluation",
			})
		}
	}

	return report
}

// ImportMetricScores converts a slice of MetricScores into a Rubric report.
func ImportMetricScores(scores []llmops.MetricScore, opts ...ImportOptions) *rubric.Rubric {
	result := &llmops.EvalResult{Scores: scores}
	return ImportEvalResult(result, opts...)
}

// MetricScoreToCategory converts a single MetricScore to a CategoryResult.
func MetricScoreToCategory(score llmops.MetricScore) rubric.CategoryResult {
	numericScore := denormalizeScore(score.Score)
	return rubric.CategoryResult{
		Category:     score.Name,
		Score:        numericScoreToScoreValue(numericScore),
		NumericScore: &numericScore,
		Reasoning:    score.Reason,
	}
}

// numericScoreToScoreValue converts a numeric score (0-10) to a categorical ScoreValue.
// Score >= 7.0 -> pass, 5.0-7.0 -> partial, < 5.0 -> fail
func numericScoreToScoreValue(score float64) rubric.ScoreValue {
	if score >= 7.0 {
		return rubric.ScorePass
	}
	if score >= 5.0 {
		return rubric.ScorePartial
	}
	return rubric.ScoreFail
}

// denormalizeScore converts a 0-1 score to 0-10 range.
func denormalizeScore(score float64) float64 {
	return score * 10.0
}

// AnnotationToFinding converts an llmops Annotation to a rubric Finding.
func AnnotationToFinding(ann llmops.Annotation) rubric.Finding {
	severity := rubric.SeverityInfo
	if ann.Label != "" {
		switch ann.Label {
		case "critical", "CRITICAL":
			severity = rubric.SeverityCritical
		case "high", "HIGH":
			severity = rubric.SeverityHigh
		case "medium", "MEDIUM":
			severity = rubric.SeverityMedium
		case "low", "LOW":
			severity = rubric.SeverityLow
		}
	}

	category := ""
	if ann.Metadata != nil {
		if cat, ok := ann.Metadata["category"].(string); ok {
			category = cat
		}
	}

	return rubric.Finding{
		Severity:    severity,
		Category:    category,
		Title:       ann.Name,
		Description: ann.Explanation,
	}
}
