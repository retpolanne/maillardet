package tex

var bookReference = &ReferencedContent{}
var websiteReference = &ReferencedContent{}

var expectedBookReference = `
@book{ibge1993,
	Address = {Rio de Janeiro},
	Author = {IBGE},
	Date-Added = {2013-08-21 13:56:10 +0000},
	Date-Modified = {2013-08-21 13:56:10 +0000},
	Edition = {3},
	Organization = {http://biblioteca.ibge.gov.br/visualizacao/livros/liv23907.pdf},
	Publisher = {Centro de Documenta{\c c}{\~a}o e Dissemina{\c c}{\~a}o de Informa{\c c}{\~o}es. Funda{\c c}{\~a}o Intituto Brasileiro de Geografia e Estat{\'\i}stica},
	Title = {Normas de apresenta{\c c}{\~a}o tabular},
	Urlaccessdate = {21 ago 2013},
	Year = {1993}}
`

var expectedWebsiteReference = `
@manual{abntex2modelo-relatorio,
	Author = {Lauro C{\'e}sar Araujo},
	Date-Added = {2013-01-15 00:05:34 +0000},
	Date-Modified = {2015-04-27 22:43:18 +0000},
	Organization = {Equipe abnTeX2},
	Title = {Modelo Can{\^o}nico de Relat{\'o}rio T{\'e}cnico e/ou Cient{\'\i}fico com abnTeX2},
	Url = {http://www.abntex.net.br/},
	Year = {2015},
	Bdsk-Url-1 = {http://www.abntex.net.br/}}
`

var providedAccentuation = `áàâãéêèîíôõóûúç`
var expectedAccentuation = `
{\'a}{\` + "`" + `a}{\^a}{\~a}{\'e}{\^e}{\` + "`" + `e}
{\^\i}{\'\i}{\~o}{\^o}{\'o}{\^u}{\'u}{\c c}
`
