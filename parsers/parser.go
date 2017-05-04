package parsers

type ParserResult interface {
	Name() string
	Raw() string
}

type Parser func(lines []string) (ParserResult, []string)

var AllParser = NewMultiParser(
	[]Parser{
		ParseContract,
		ParseAssets,
		ParseCargoScan,
		ParseDScan,
	})

// ParseAssets (type func([]string) (ParserResult, []string)) as type func([]string) ([]ParserResult, []string) in field value

type MultiParserResult struct {
	results []ParserResult
}

func (r *MultiParserResult) Name() string {
	return "multi"
}

func (r *MultiParserResult) Raw() string {
	// TODO: append the Raw() of every parser result
	return ""
}

func NewMultiParser(parsers []Parser) Parser {
	return Parser(
		func(lines []string) (ParserResult, []string) {
			multiParserResult := &MultiParserResult{
				results: make([]ParserResult, 0),
			}
			left := lines
			for _, parser := range parsers {
				var result ParserResult
				result, left = parser(left)
				if result != nil {
					multiParserResult.results = append(multiParserResult.results, result)
				}
				if len(left) == 0 {
					break
				}
			}
			return multiParserResult, left
		})
}