package parser

type ParserType int

const (
	ParserType_None             = 0   // none
	ParserType_AndroidWandoujia = 1   // 豌豆荚
	ParserType_AndroidAnzhi     = 2   // 安智
	ParserType_IosAppStore      = 100 // 苹果应用商店
	ParserType_IosAnn9          = 101 // ann9爬的ios网站
	ParserTypeApe51             = 200 // 51ape网站
	ParserType_FileXuexi111     = 201 // 爬文件
	ParserType_FileDowncc       = 202
	ParserType_FileGdajie       = 203
	ParserType_FileJava1234     = 204
	ParserType_FilePdfzj        = 205
)
