package parser

type ParserType int

const (
	ParserTypeWandoujia = iota // 0 豌豆荚
	ParserTypeAppStore         // 苹果应用商店
	ParserTypeApe51            // 51ape网站
	ParserTypeAnn9             // ann9爬的ios网站
)
