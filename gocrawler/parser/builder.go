package parser

// NewParser 新建解析器
func NewParser(parserType ParserType) IParser {
	var p IParser
	switch parserType {
	case ParserTypeAppStore:
		p = new(AppStoreParser)
	case ParserTypeWandoujia:
		p = new(WandoujiaParser)
	case ParserTypeApe51:
		p = new(Ape51Parser)
	case ParserTypeAnn9:
		p = new(Ann9Parser)
	case ParserTypeFile:
		p = new(FileParser)
	default:
	}

	if p == nil {
		return nil
	}
	p.SetupData()

	return p
}
