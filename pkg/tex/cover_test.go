package tex

var cover = &Cover{}

var expectedCover = `
\titulo{Título do artigo}
\autor{Equipe \abnTeX}
\local{Brasil}
\data{2018, v-1.9.7}
\orientador{João da Silva}
\coorientador{Equipe \abnTeX}
\instituicao{%
  Universidade do Brasil -- UBr
  \par
  Faculdade de Arquitetura da Informação
  \par
  Programa de Pós-Graduação}
\tipotrabalho{Tese (Doutorado)}
\preambulo{Modelo canônico de trabalho monográfico acadêmico em conformidade com
as normas ABNT apresentado à comunidade de usuários \LaTeX.}
`

var expectedTitle = `Modelo Canônico de\\ Trabalho Acadêmico com \abnTeX`
