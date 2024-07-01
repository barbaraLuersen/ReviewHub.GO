package main

import (
	"fmt"
	"reviewhub/database"
	"reviewhub/model"
	"reviewhub/routes"
)

func main() {
	model.Series = []model.Serie{
		{
			Id: 1, Nome: "This is Us", Datalancamento: "2016-09-20", Temporadas: 6,
			Resenha: "This is Us é uma série emocionante que segue a vida de três irmãos e seus pais. A série aborda temas complexos como adoção, identidade, saúde mental e família, utilizando uma narrativa não linear que revela segredos e conexões inesperadas ao longo do tempo. Com personagens profundamente desenvolvidos e atuações tocantes, 'This is Us' cativa o público com suas histórias comoventes e autênticas.",
		},
		{
			Id: 2, Nome: "Modern Family", Datalancamento: "2009-09-23", Temporadas: 11,
			Resenha: "Modern Family é uma série de comédia que acompanha as vidas de três famílias interligadas, oferecendo uma visão divertida e sincera das complexidades da vida familiar moderna. Com um estilo de documentário fictício, a série explora as dinâmicas",
		},
		{
			Id: 3, Nome: "The Office", Datalancamento: "2005-03-24", Temporadas: 9,
			Resenha: "The Office é uma série de comédia no formato de documentário fictício que retrata o dia a dia dos funcionários da Dunder Mifflin, uma empresa de papel em Scranton, Pensilvânia. Com um elenco memorável liderado por Steve Carell como o peculiar chefe Michael Scott, a série combina humor, situações absurdas e momentos tocantes, capturando a essência do ambiente de trabalho com uma mistura de sátira e realismo.",
		},
		{
			Id: 4, Nome: "Fleabag", Datalancamento: "2016-07-21", Temporadas: 2,
			Resenha: "Fleabag é uma série de comédia dramática britânica criada e estrelada por Phoebe Waller-Bridge. A série segue a vida de uma jovem mulher em Londres, apelidada de Fleabag, que navega pelas complexidades de relacionamentos, família e autoaceitação. Conhecida por sua escrita afiada, humor negro e quebras da quarta parede, Fleabag oferece uma perspectiva crua e hilária sobre as lutas e triunfos da vida moderna.",
		},
	}
	database.ConexaoDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
