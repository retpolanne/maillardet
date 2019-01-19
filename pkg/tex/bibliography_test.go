package tex

var bookReference = &ReferencedContent{}
var websiteReference = &ReferencedContent{}

var expectedBookReference = ``

var expectedWebsiteReference = ``

var providedAccentuation = `áàâãéêèîíôõóûúç`
var expectedAccentuation = `
{\'a}{\` + "`" + `a}{\^a}{\~a}{\'e}{\^e}{\` + "`" + `e}
{\^\i}{\'\i}{\~o}{\^o}{\'o}{\^u}{\'u}{\c c}
`
