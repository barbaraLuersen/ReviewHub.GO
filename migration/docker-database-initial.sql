create table series(
    id serial primary key,
    nome varchar,
    dataLancamento varchar,
    temporadas int,
    resenha varchar
);

INSERT INTO series(nome, dataLancamento, temporadas, resenha) VALUES
('This is Us', '20-09-2016', 6, 'This is Us é uma série emocionante que segue a vida de três irmãos e seus pais. A série aborda temas complexos como adoção, identidade, saúde mental e família, utilizando uma narrativa não linear que revela segredos e conexões inesperadas ao longo do tempo. Com personagens profundamente desenvolvidos e atuações tocantes, This is Us cativa o público com suas histórias comoventes e autênticas.'),
('Modern Family','23-09-2009', 11, 'Modern Family é uma série de comédia que acompanha as vidas de três famílias interligadas, oferecendo uma visão divertida e sincera das complexidades da vida familiar moderna. Com um estilo de documentário fictício, a série explora as dinâmicas'),
('The Office', '24-03-2005', 9, 'The Office é uma série de comédia no formato de documentário fictício que retrata o dia a dia dos funcionários da Dunder Mifflin, uma empresa de papel em Scranton, Pensilvânia. Com um elenco memorável liderado por Steve Carell como o peculiar chefe Michael Scott, a série combina humor, situações absurdas e momentos tocantes, capturando a essência do ambiente de trabalho com uma mistura de sátira e realismo.'),
('Fleabag', '21-07-2016', 2, 'Fleabag é uma série de comédia dramática britânica criada e estrelada por Phoebe Waller-Bridge. A série segue a vida de uma jovem mulher em Londres, apelidada de Fleabag, que navega pelas complexidades de relacionamentos, família e autoaceitação. Conhecida por sua escrita afiada, humor negro e quebras da quarta parede, Fleabag oferece uma perspectiva crua e hilária sobre as lutas e triunfos da vida moderna.');