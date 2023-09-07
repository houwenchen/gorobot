package lexer

type Settings struct {
	Names                    []string
	Aliases                  map[string]string
	MultiUse                 []string
	SingleValue              []string
	NameAndArguments         []string
	NameArgumentsAndWithName []string

	SettingsMap map[string]StatementTokens
}

type FileSettings struct {
	Settings
}

type SuiteFileSettings struct {
	FileSettings

	Names   []string
	Aliases map[string]string
}

type InitFileSettings struct {
	FileSettings

	Names   []string
	Aliases map[string]string
}

type ResourceFileSettings struct {
	FileSettings

	Names []string
}

type TestCaseSettings struct {
	Settings

	Names []string
}

type KeywordSettings struct {
	Settings

	Names []string
}

func NewSettings() *Settings {
	s := &Settings{
		Names:   make([]string, 0),
		Aliases: make(map[string]string),
		MultiUse: []string{
			"Metadata",
			"Library",
			"Resource",
			"Variables"},
		SingleValue: []string{
			"Resource",
			"Test Timeout",
			"Test Template",
			"Timeout",
			"Template",
			"Name",
		},
		NameAndArguments: []string{
			"Metadata",
			"Suite Setup",
			"Suite Teardown",
			"Test Setup",
			"Test Teardown",
			"Test Template",
			"Setup",
			"Teardown",
			"Template",
			"Resource",
			"Variables",
		},
		NameArgumentsAndWithName: []string{
			"Library",
		},
	}

	settingsMap := make(map[string]StatementTokens)
	for _, obj := range s.Names {
		settingsMap[obj] = nil
	}

	s.SettingsMap = settingsMap

	return s
}

func (s *Settings) lex(statement StatementTokens) {
	// orig = s._format_name(statement[0].value)
}

// TODO: 看着没什么作用，后续评估是否需要留着
func (s *Settings) _format_name(name string) string {
	return name
}

func (s *Settings) _validate(orig string, name string, statement StatementTokens) {

}

func NewSuiteFileSettings() *SuiteFileSettings {
	return &SuiteFileSettings{
		Names: []string{
			"Documentation",
			"Metadata",
			"Name",
			"Suite Setup",
			"Suite Teardown",
			"Test Setup",
			"Test Teardown",
			"Test Template",
			"Test Timeout",
			"Test Tags",
			"Default Tags",
			"Keyword Tags",
			"Library",
			"Resource",
			"Variables",
		},
		Aliases: map[string]string{
			"Force Tags":    "Test Tags",
			"Task Tags":     "Test Tags",
			"Task Setup":    "Test Setup",
			"Task Teardown": "Test Teardown",
			"Task Template": "Test Template",
			"Task Timeout":  "Test Timeout",
		},
	}
}

func NewInitFileSettings() *InitFileSettings {
	return &InitFileSettings{
		Names: []string{
			"Documentation",
			"Metadata",
			"Name",
			"Suite Setup",
			"Suite Teardown",
			"Test Setup",
			"Test Teardown",
			"Test Timeout",
			"Test Tags",
			"Keyword Tags",
			"Library",
			"Resource",
			"Variables",
		},
		Aliases: map[string]string{
			"Force Tags":    "Test Tags",
			"Task Tags":     "Test Tags",
			"Task Setup":    "Test Setup",
			"Task Teardown": "Test Teardown",
			"Task Timeout":  "Test Timeout",
		},
	}
}

func NewResourceFileSettings() *ResourceFileSettings {
	return &ResourceFileSettings{
		Names: []string{
			"Documentation",
			"Keyword Tags",
			"Library",
			"Resource",
			"Variables",
		},
	}
}

func NewTestCaseSettings() *TestCaseSettings {
	return &TestCaseSettings{
		Names: []string{
			"Documentation",
			"Tags",
			"Setup",
			"Teardown",
			"Template",
			"Timeout",
		},
	}
}

func NewKeywordSettings() *KeywordSettings {
	return &KeywordSettings{
		Names: []string{
			"Documentation",
			"Arguments",
			"Teardown",
			"Timeout",
			"Tags",
			"Return",
		},
	}
}
