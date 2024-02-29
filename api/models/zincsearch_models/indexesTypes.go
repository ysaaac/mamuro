package zincsearch_models

import (
	"api/config"
	"strconv"
)

// There's two possibles approaches of these functions. To create totally independent functions to declare structures
// or create one as a base and the others just modify the created structures as needed

func defaultTypeData() CreateIndex {
	// zincShardNum is the value of CPU cores it will use to improve performance.
	// I'm setting it dynamically, because it will depend on the CPU I'm running it.
	zincShardNum, _ := strconv.Atoi(config.GetEnv("ZINC_SHARD_NUM", "1"))

	defaultSettings := IndexSettings{
		Analysis: Analysis{
			Analyzer: Analyzer{
				Default: DefaultAnalyzer{
					Type: "standard",
				},
				SymbolicAnalyzer: ExtraAnalyzerConfig{
					Tokenizer:  "standard",
					CharFilter: []string{"unicode_mapping"},
				},
			},
			CharFilter: CharFilter{
				UnicodeMapping: Filter{
					Type:     "mapping",
					Mappings: []string{"\\u003c => <", "\\u003e => >"},
				},
			},
		},
	}

	return CreateIndex{
		StorageType: "disk",
		ShardNum:    zincShardNum,
		Settings:    defaultSettings,
	}
}

func InboxIndexTypeStructure() CreateIndex {
	inboxStructure := defaultTypeData()
	inboxStructure.Name = "inbox"
	inboxStructure.Mappings = IndexMappings{
		Properties: map[string]IndexProperty{
			"message_id": {
				Type:  "keyword",
				Index: true,
			},
			"date": {
				Type:     "date",
				Index:    true,
				Sortable: true,
				Format:   "time.RFC1123",
			},
			"from": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"to": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"subject": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"mime_version": {
				Type: "text",
			},
			"content_type": {
				Type: "text",
			},
			"content_transfer_encoding": {
				Type: "text",
			},
			"x_from": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"x_to": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"x_cc": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"x_bcc": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"x_folder": {
				Type: "text",
			},
			"x_origin": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"x_file_name": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
			"content": {
				Type:          "text",
				Index:         true,
				Highlightable: true,
			},
		},
	}

	return inboxStructure
}

func SentItemsIndexTypeStructure() CreateIndex {
	// For now there's no difference between Inbox fields and Sent Items fields, so I'm reusing the Inbox, but
	// probably could be different in the future
	sentItemsStructure := InboxIndexTypeStructure()
	sentItemsStructure.Name = "sent_items"
	return sentItemsStructure
}
