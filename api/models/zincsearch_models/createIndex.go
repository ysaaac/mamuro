package zincsearch_models

type CreateIndex struct {
	Name        string        `json:"name"`
	StorageType string        `json:"storage_type"`
	ShardNum    int           `json:"shard_num"`
	Mappings    IndexMappings `json:"mappings"`
	Settings    IndexSettings `json:"settings"`
}

type IndexMappings struct {
	Properties map[string]IndexProperty `json:"properties"`
}

type IndexProperty struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store,omitempty"`
	Highlightable bool   `json:"highlightable,omitempty"`
	Format        string `json:"format,omitempty"`
	Sortable      bool   `json:"sortable,omitempty"`
	Aggregatable  bool   `json:"aggregatable,omitempty"`
}

type IndexSettings struct {
	Analysis Analysis `json:"analysis"`
}

type Analysis struct {
	Analyzer   Analyzer   `json:"analyzer"`
	CharFilter CharFilter `json:"char_filter,omitempty"`
}

type Analyzer struct {
	Default          DefaultAnalyzer     `json:"default"`
	SymbolicAnalyzer ExtraAnalyzerConfig `json:"symbolic_analyzer,omitempty"`
}

type DefaultAnalyzer struct {
	Type string `json:"type"`
}

type ExtraAnalyzerConfig struct {
	Tokenizer  string   `json:"tokenizer"`
	CharFilter []string `json:"char_filter"`
}
type CharFilter struct {
	UnicodeMapping Filter `json:"unicode_mapping"`
}

type Filter struct {
	Type     string   `json:"type"`
	Mappings []string `json:"mappings"`
}
