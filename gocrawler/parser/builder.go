package parser

// NewParser 新建解析器
func NewParser(parserType ParserType) IParser {
	var p IParser
	switch parserType {
	case ParserType_AndroidWandoujia:
		p = new(AndroidWandoujiaParser)
	case ParserType_AndroidAnzhi:
		p = new(AndroidAnzhiParser)
	case ParserType_IosAppStore:
		p = new(AppStoreParser)
	case ParserType_IosAnn9:
		p = new(Ann9Parser)
	case ParserTypeApe51:
		p = new(Ape51Parser)
	case ParserType_FileXuexi111:
		p = new(Xuexi111FileParser)
	case ParserType_FileDowncc:
		p = new(DownccFileParser)
	case ParserType_FileGdajie:
		p = new(GdajieFileParser)
	case ParserType_FileJava1234:
		p = new(Java1234FileParser)
	case ParserType_FilePdfzj:
		p = new(PdfzjFileParser)
	default:
	}

	if p == nil {
		return nil
	}
	p.SetupData()

	return p
}
