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
	Publisher = {Centro de Documentação e Disseminação de Informações. Fundação Instituto Brasileiro de Geografia e Estatística},
	Title = {Normas de apresentação tabular},
	Urlaccessdate = {21 ago 2013},
	Year = {1993}}
`

var expectedWebsiteReference = `
@manual{babel,
	Author = {Johannes Braams},
	Date-Added = {2013-02-17 13:37:14 +0000},
	Date-Modified = {2013-02-17 13:38:38 +0000},
	Month = {Apr.},
	Title = {Babel, a multilingual package for use with LATEX's standard document classes},
	Url = {http://mirrors.ctan.org/info/babel/babel.pdf},
	Urlaccessdate = {17 fev. 2013},
	Year = {2008},
	Bdsk-Url-1 = {http://mirrors.ctan.org/info/babel/babel.pdf}}
`
