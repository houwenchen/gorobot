package builder

type TestSuiteBuilder struct {
	StandardParsers    map[string]Parser
	CustomParsers      map[string]CustomParser
	Defaults           TestDefaults
	IncludedExtensions []string
	IncludedFiles      []string
	Rpa                bool
	AllowEmptySuite    bool
}

func NewTestSuiteBuilder(
	included_extensions []string,
	included_files []string,
	custom_parsers CustomParser,
	defaults TestDefaults,
	rpa, allow_empty_suite bool,
) *TestSuiteBuilder {
	b := &TestSuiteBuilder{
		Defaults:        defaults,
		IncludedFiles:   included_files,
		Rpa:             rpa,
		AllowEmptySuite: allow_empty_suite,
	}

	b.StandardParsers = b.get_standard_parsers()
	b.CustomParsers = b.get_custom_parsers()
	b.IncludedExtensions = b.get_included_extensions(included_extensions)

	return b
}

func (b *TestSuiteBuilder) get_standard_parsers() map[string]Parser {
	// robot_parser := NewRobotParser()
	// rest_parser := NewRestParser()
	// json_parser := NewJsonParser()

	return nil
}

func (b *TestSuiteBuilder) get_custom_parsers() map[string]CustomParser {
	return nil
}

func (b *TestSuiteBuilder) get_included_extensions(included_extensions []string) []string {
	if len(included_extensions) == 0 {
		return []string{".gorobot"}
	}

	return included_extensions
}
