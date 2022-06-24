package hover

import (
	"reflect"
	"testing"

	sglsp "github.com/sourcegraph/go-lsp"
	"github.com/stretchr/testify/assert"

	"github.com/snyk/snyk-ls/domain/snyk/issues"
	"github.com/snyk/snyk-ls/internal/observability/ux"
	"github.com/snyk/snyk-ls/internal/uri"
)

func setupFakeHover() sglsp.DocumentURI {
	target := NewDefaultService(ux.NewTestAnalytics()).(*DefaultHoverService)
	fakeHover := []Hover[Context]{
		{Range: sglsp.Range{
			Start: sglsp.Position{Line: 3, Character: 56},
			End:   sglsp.Position{Line: 5, Character: 80},
		},
		},
	}

	filePath := uri.PathToUri("file:///fake-file.txt")
	target.hovers[filePath] = fakeHover
	target.hoverIndexes[uri.PathFromUri(filePath+"rangepositionstuff")] = true

	return filePath
}

func Test_registerHovers(t *testing.T) {
	target := NewDefaultService(ux.NewTestAnalytics()).(*DefaultHoverService)
	documentUri := uri.PathToUri("fake-file.json")
	hover := DocumentHovers{
		Uri: documentUri,
		Hover: []Hover[Context]{
			{
				Id: "test-id",
				Range: sglsp.Range{
					Start: sglsp.Position{
						Line:      10,
						Character: 14,
					},
					End: sglsp.Position{
						Line:      56,
						Character: 87,
					},
				},
				Message: "Very important hover",
			},
		},
	}

	target.registerHovers(hover)
	// assert de-duplication
	target.registerHovers(hover)

	assert.Equal(t, len(target.hovers[documentUri]), 1)
	assert.Equal(t, len(target.hoverIndexes), 1)
}

func Test_DeleteHover(t *testing.T) {
	target := NewDefaultService(ux.NewTestAnalytics()).(*DefaultHoverService)
	documentUri := setupFakeHover()
	target.DeleteHover(documentUri)

	assert.Equal(t, len(target.hovers[documentUri]), 0)
	assert.Equal(t, len(target.hoverIndexes), 0)
}

func Test_ClearAllHovers(t *testing.T) {
	target := NewDefaultService(ux.NewTestAnalytics()).(*DefaultHoverService)
	documentUri := setupFakeHover()
	target.ClearAllHovers()

	assert.Equal(t, len(target.hovers[documentUri]), 0)
	assert.Equal(t, len(target.hoverIndexes), 0)
}

func Test_GetHoverMultiline(t *testing.T) {
	target := NewDefaultService(ux.NewTestAnalytics()).(*DefaultHoverService)

	tests := []struct {
		hoverDetails []Hover[Context]
		query        sglsp.Position
		expected     Result
	}{
		// multiline range
		{
			hoverDetails: []Hover[Context]{{Range: sglsp.Range{
				Start: sglsp.Position{Line: 3, Character: 56},
				End:   sglsp.Position{Line: 5, Character: 80},
			},
				Message: "## Vulnerabilities found"}},
			query: sglsp.Position{Line: 4, Character: 66},
			expected: Result{Contents: MarkupContent{
				Kind: "markdown", Value: "## Vulnerabilities found"},
			},
		},
		// exact line but within character range
		{
			hoverDetails: []Hover[Context]{{Range: sglsp.Range{
				Start: sglsp.Position{Line: 4, Character: 56},
				End:   sglsp.Position{Line: 4, Character: 80},
			},
				Message: "## Vulnerabilities found"}},
			query: sglsp.Position{Line: 4, Character: 66},
			expected: Result{Contents: MarkupContent{
				Kind: "markdown", Value: "## Vulnerabilities found"},
			},
		},
		// exact line and exact character
		{
			hoverDetails: []Hover[Context]{{Range: sglsp.Range{
				Start: sglsp.Position{Line: 4, Character: 56},
				End:   sglsp.Position{Line: 4, Character: 56},
			},
				Message: "## Vulnerabilities found"}},
			query: sglsp.Position{Line: 4, Character: 56},
			expected: Result{Contents: MarkupContent{
				Kind: "markdown", Value: "## Vulnerabilities found"},
			},
		},
		// hover left of the character position on exact line
		{
			hoverDetails: []Hover[Context]{{Range: sglsp.Range{
				Start: sglsp.Position{Line: 4, Character: 56},
				End:   sglsp.Position{Line: 4, Character: 86},
			},
				Message: "## Vulnerabilities found"}},
			query: sglsp.Position{Line: 4, Character: 45},
			expected: Result{Contents: MarkupContent{
				Kind: "markdown", Value: ""},
			},
		},
		// hover right of the character position on exact line
		{
			hoverDetails: []Hover[Context]{{Range: sglsp.Range{
				Start: sglsp.Position{Line: 4, Character: 56},
				End:   sglsp.Position{Line: 4, Character: 86},
			},
				Message: "## Vulnerabilities found"}},
			query: sglsp.Position{Line: 4, Character: 105},
			expected: Result{Contents: MarkupContent{
				Kind: "markdown", Value: ""},
			},
		},
	}

	path := uri.PathToUri("path/to/package.json")
	for _, tc := range tests {
		target.ClearAllHovers()
		target.hovers[path] = tc.hoverDetails

		result := target.GetHover(path, tc.query)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("expected: %v, got: %v", tc.expected, result)
		}
	}
}

func Test_TracksAnalytics(t *testing.T) {
	analytics := ux.NewTestAnalytics()
	target := NewDefaultService(analytics).(*DefaultHoverService)

	path := uri.PathToUri("path/to/package.json")
	target.ClearAllHovers()
	target.hovers[path] = []Hover[Context]{
		{
			Context: issues.Issue{
				ID:        "issue",
				Severity:  issues.Medium,
				IssueType: issues.ContainerVulnerability,
			},
			Range: sglsp.Range{
				Start: sglsp.Position{Line: 3, Character: 56},
				End:   sglsp.Position{Line: 5, Character: 80},
			},
			Message: "## Vulnerabilities found"},
	}

	target.GetHover(path, sglsp.Position{Line: 4, Character: 66})
	assert.Len(t, analytics.GetAnalytics(), 1)
	assert.Equal(t, ux.IssueHoverIsDisplayedProperties{
		IssueId:   "issue",
		IssueType: ux.ContainerVulnerability,
		Severity:  ux.Medium,
	}, analytics.GetAnalytics()[0])
}
