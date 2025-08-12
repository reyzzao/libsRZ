
---

### LOGICA FUNCAO LOADER_INDEX VIA RUNTIME
- explicacao para implemenatcao no Runtime : TS Deno [by Gmn] (https://g.co/gemini/share/e940305000f7)

### Toda_Explicacao
- É uma ótima pergunta! Entender a lógica por trás de um loader como esse permite que você aplique o mesmo conceito em qualquer outra linguagem ou ambiente, como Python, Node.js, Go, etc.
A lógica da função loaderModulesIndex pode ser resumida em quatro passos conceituais que são universais para qualquer runtime que permita carregamento dinâmico de código.

- Passos
  1. Iterar sobre os Alvos
  A função começa com um laço for que percorre um array de strings (moduleFolders). Este passo é o "motor" do nosso loader. Ele garante que a lógica será executada uma vez para cada pasta de módulo que você quer carregar.
  
  Na função: for (const folderName of moduleFolders)
  
  Conceito Universal: Qualquer linguagem usa um laço (for, foreach, map) para processar uma coleção de itens, um por vez.
  
  2. Construir Dinamicamente o Caminho do Arquivo
  Este é o passo mais crucial. O código não tem o caminho para o arquivo Index.ts "fixo". Em vez disso, ele o constroi em tempo de execução, juntando o caminho base (basePath) com o nome da pasta e o nome do arquivo-alvo (Index.ts).
  
  Na função: new URL(./${folderName}/Index.ts, basePath).href

  Conceito Universal: Usar funções do runtime para juntar caminhos (os.path.join() em Python, path.join() em Node.js) em vez de concatenar strings manualmente. O objetivo é sempre criar um caminho de arquivo válido para o sistema.

  3. Carregar o Módulo de Forma Assíncrona
  Aqui o loader executa sua função principal. Usando o caminho de arquivo que foi construído no passo 2, o runtime é instruído a carregar e executar o código contido naquele arquivo. Isso é feito de forma assíncrona para não travar a aplicação, pois a leitura do disco pode levar tempo.

  Na função: await import(filePath)

  Conceito Universal: A maioria dos runtimes modernos tem uma forma de carregar módulos dinamicamente. Em Node.js, é com import() ou require() dentro de um escopo dinâmico. Em Python, é com a biblioteca importlib. O conceito é sempre "diga-me o caminho, e eu te darei o módulo".

  4. Coletar e Agregar as Exportações
  Este é o passo final. Cada import() retorna as exportações do módulo de forma separada. O loader precisa juntar todas essas exportações de volta em um único objeto para que o seu código principal (Init.ts) possa acessá-las de forma unificada.

  Na função: Object.assign(exports_collection, module)

  Conceito Universal: Combinar os resultados de cada iteração em uma estrutura de dados final. Isso pode ser feito mesclando objetos (como o Object.assign faz), juntando arrays, ou adicionando chaves a um dicionário.

- Resumo da Lógica Universal:
  Você pode pensar no padrão de um loader assim:
  Para cada (módulo) em (lista_de_módulos):
  Construa o caminho do arquivo de entrada do módulo.
  Carregue o conteúdo desse arquivo dinamicamente.
  Junte o conteúdo carregado a uma coleção central.
  O resultado é um objeto final que centraliza todo o código de todos os seus módulos.

---