# Exercícios em Go

Este repositório contém exercícios práticos de programação em Go.

## Tutorial - Como compilar e executar

### Opção 1: Compilar e executar diretamente
```bash
go run file.go
```

### Opção 2: Criar executável e depois executar
```bash
# Compilar e criar executável
go build file.go

# Executar o arquivo gerado
./file.exe
```

Semana 2, dia 22/08, primeira aula

Discutimos sobre os alunos que não estão mais em aula.

Os alunos escolheram ter uma prova de recuperação no último dia letivo, 19 de dezembro. O professor escolheu como primeira nota um portfólio.

A segunda nota foi escolhida por nós, e será um projeto feito em grupos de 4 pessoas ou individualmente. A apresentação do projeto será nos dias 5 e 12 de dezembro. Nenhuma das notas terá peso.

O portfólio pode ser feito usando qualquer técnica, ferramenta e tecnologia. Um documento, página, etc.

No projeto não será permitido o uso de nenhum LLM (agente de IA). Se houver indícios de uso, a nota será zerada.

O professor também optou por termos listas de exercícios durante o semestre, nas horas fora de sala. Será um trabalho de programação concorrente usando a linguagem GO ou GOLANG. O professor irá dar uma aula rápida de GO para entendermos os seus tipos, fundamentos, fluxo de controle, estruturas de dados e integração com banco de dados.

Em um segundo momento, ele nos explicará os fundamentos de concorrência, paralelismo, e padrões de concorrências mais básicos.

Em um terceiro momento, veremos um framework para nós auxiliar na programação de concorrência. Também sincronização, tratamentos de erros e ferramentas.

No módulo quatro, veremos exemplos de aplicações práticas.

Meu grupo do projeto terá o João Lucas, Samara e Davi.

Os alunos negociaram 0,5 ponto bônus na média final para quem tiver 95% de presença até o final do semestre (5% equivale a um dia de aula, com 3 aulas de 45 minutos), além de outro 0,5 ponto para aqueles que vieram no primeiro dia de aula (22/08).

Por meio de uma votação, o professor nos fez decidir a quantidade de exercícios. As opções eram 10, o dobro de 10, ou o dobro do anterior dividido por três, e escolhemos 10 por semana, e começaremos a partir da semana que vem.


Semana 3
Dia 29/08
O professor nos passou 25 exercícios de algoritmo básicos para praticarmos golang.
São eles:

1 - soma de dois números inteiros

2 - divisão de dois números inteiros

3 - um algoritmo que recebe um número inteiro e imprime o antecessor e sucessor

4 - verificar se um número é par, positivo ou negativo

5 - verificar se um número é primo ou não

6 - que faça a ordenação de uma sequencia numérica

7 - que faça a ordenação de caracteres em ordem ascendente

8 - que implemente uma arvore de decisão binaria

9 - imprime o valor e endereço de uma variável

10 - que resolva a brincadeira de peças da torre de Hanói de três discos

11  - que recebe a data de nascimento e retorne o dia da semana

12 - que retorne um booleano e a igualdade entre dois números

13 - que retorne a moda de uma sequencia numérica

14 - verifique se uma sequencia de caracteres formam um palíndromo

15 - calcule a área de um retângulo

16 - que calcule conversão entre unidades de temperatura

17 - que simule o jogo da adivinhação

18 - que receba 3 matrizes de inteiros sem sinal de 8 bits e grave um arquivo .jpg

19 - que conte a quantidade de vogais e consoantes presentes em um texto

20 - que encontre padrões de ocorrências de palavras em um texto

21 - que calcule o fatorial de um número n

22 - que imprima Hello World!

23 - que calcule IMC padrão Brasil

24 - que calcule MMC

25 - que calcule a média de 2 ou mais números


Semana 4
Dia 03/09
Primeiramente fiz o exercício 22 para entender a estrutura da linguagem, e depois os 5 primeiros exercícios.

Dia 05/09
Características principais da GOLANG
	Compilada - foco em alto desempenho
	Maior segurança, ao compilar antes, previne certos erros
	Concorrência, suporta nativamente concorrência de processos
	Tipada - Previne confusão com tipos de variáveis
	Desenvolvimento de microsserviços e em nuvem
	Foco em backend

Documentação da GOLANG - simples e fácil
	https://go.dev/doc/ - oficial
	https://gobyxample.com - contém exemplos e tutoriais


1 - Concorrência
2 - Simplicidade e desempenho - Eu, João Lucas e Victor Cazuo
3 - Compilação e tipagem
4 - Coletor de lixo *ou* Garbage Collector
5 - Biblioteca padrão

Dentro do meu grupo:
	Linguagem preferida: Python
	Considerando nossa linguagem favorita, escolher um representante - João Lucas - para defender GO em relação ao python, levando em conta a nossa característica escolhida.

Argumentação:
 - GO é compilada diretamente para código de máquina, enquanto o python necessita de um interpretador, o que a torna GOLANG mais rápida, extremamente eficiente e menos suscetível à erros, além de simplificar o deploy e reduzir as dependências;
- GO aproveita o potencial de múltiplos núcleos de processamento;
- Em Go, a organização em pacotes importáveis deixa o código mais limpo e gerenciável;


Programação em python possui uma estrutura mais livre em comparação a GO, que possui uma estrutura pré-definida, ou seja, Python necessita de menos código;
Python pode ser aplicado em várias áreas.

Escolha GO para performance bruta, eficiência e foco em microsserviços;
	Escolha Python para projetos de rápido desenvolvimento, facilidade de uso e acesso a um rico ecossistema

Dia 17/10
O professor fez o ex de ppm, falar sobre

Dia 24/10
O professor revisou atributos, falar sobre
tarefa: descobrir diferença entre "defer" e "panick"