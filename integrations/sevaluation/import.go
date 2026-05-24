package sevaluation

import (
	"github.com/plexusone/omniobserve/llmops"
	"github.com/plexusone/structured-evaluation/evaluation"
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
	// Default: evaluation.DefaultPassCriteria()
	PassCriteria evaluation.PassCriteria
}

// DefaultImportOptions returns the default import configuration.
func DefaultImportOptions() ImportOptions {
	return ImportOptions{
		ReviewType:   "llm_evaluation",
		Document:     "",
		PassCriteria: evaluation.DefaultPassCriteria(),
	}
}

// ImportEvalResult converts an llmops EvalResult into an EvaluationReport.
// Each MetricScore becomes a CategoryResult in the report.
func ImportEvalResult(result *llmops.EvalResult, opts ...ImportOptions) *evaluation.EvaluationReport {
	opt := DefaultImportOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}

	report := evaluation.NewEvaluationReport(opt.ReviewType, opt.Document)
	report.PassCriteria = opt.PassCriteria

	for _, score := range result.Scores {
		numericScore := denormalizeScore(score.Score)
		cat := evaluation.CategoryResult{
			Category:     score.Name,
			Score:        numericScoreToScoreValue(numericScore),
			NumericScore: &numericScore,
			Reasoning:    score.Reason,
		}
		report.AddCategoryResult(cat)

		// If score has an error, add it as a finding
		if score.Error != "" {
			report.AddFinding(evaluation.Finding{
				Severity:       evaluation.SeverityMedium,
				Category:       score.Name,
				Title:          "Evaluation error",
				Description:    score.Error,
				Recommendation: "Review and re-run evaluation",
			})
		}
	}

	return report
}

// ImportMetricScores converts a slice of MetricScores into an EvaluationReport.
func ImportMetricScores(scores []llmops.MetricScore, opts ...ImportOptions) *evaluation.EvaluationReport {
	result := &llmops.EvalResult{Scores: scores}
	return ImportEvalResult(result, opts...)
}

// MetricScoreToCategory converts a single MetricScore to a CategoryResult.
func MetricScoreToCategory(score llmops.MetricScore) evaluation.CategoryResult {
	numericScore := denormalizeScore(score.Score)
	return evaluation.CategoryResult{
		Category:     score.Name,
		Score:        numericScoreToScoreValue(numericScore),
		NumericScore: &numericScore,
		Reasoning:    score.Reason,
	}
}

// numericScoreToScoreValue converts a numeric score (0-10) to a categorical ScoreValue.
// Score >= 7.0 -> pass, 5.0-7.0 -> partial, < 5.0 -> fail
func numericScoreToScoreValue(score float64) evaluation.ScoreValue {
	if score >= 7.0 {
		return evaluation.ScorePass
	}
	if score >= 5.0 {
		return evaluation.ScorePartial
	}
	return evaluation.ScoreFail
}

// denormalizeScore converts a 0-1 score to 0-10 range.
func denormalizeScore(score float64) float64 {
	return score * 10.0
}

// AnnotationToFinding converts an llmops Annotation to an evaluation Finding.
func AnnotationToFinding(ann llmops.Annotation) evaluation.Finding {
	severity := evaluation.SeverityInfo
	if ann.Label != "" {
		switch ann.Label {
		case "critical", "CRITICAL":
			severity = evaluation.SeverityCritical
		case "high", "HIGH":
			severity = evaluation.SeverityHigh
		case "medium", "MEDIUM":
			severity = evaluation.SeverityMedium
		case "low", "LOW":
			severity = evaluation.SeverityLow
		}
	}

	category := ""
	if ann.Metadata != nil {
		if cat, ok := ann.Metadata["category"].(string); ok {
			category = cat
		}
	}

	return evaluation.Finding{
		Severity:    severity,
		Category:    category,
		Title:       ann.Name,
		Description: ann.Explanation,
	}
}
