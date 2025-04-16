# 📄 CSVBuilder em Go

## 📌 Objetivo  
Criar um `CSVBuilder` em Go para gerar arquivos CSV de forma **flexível e configurável**. O objetivo é permitir que os usuários construam e salvem arquivos CSV sem precisar manipular strings manualmente.

---

## 🔹 Requisitos Funcionais

### 1. Definir o nome do arquivo  
- O usuário deve poder definir um nome para o arquivo CSV antes de salvá-lo.  
- Se não for definido, o nome padrão será `"output.csv"`.  

### 2. Definir um cabeçalho (nomes das colunas)  
- O usuário pode definir um cabeçalho com os nomes das colunas do CSV.  
- Se não for definido, o CSV **não** terá cabeçalho.  

### 3. Adicionar múltiplas linhas de dados  
- O usuário pode adicionar quantas linhas desejar, cada uma representando um registro no CSV.  
- O número de elementos em cada linha **deve ser consistente** com o cabeçalho (se houver).  

### 4. Escolher um delimitador (`,` ou `;`)  
- O usuário pode escolher entre `,` (vírgula) ou `;` (ponto e vírgula) como delimitador dos campos.  
- Se não for definido, o padrão será `,`.  

### 5. Gerar o conteúdo do CSV  
- O método `Build()` deve gerar o conteúdo CSV formatado corretamente.  

### 6. Salvar o arquivo no sistema  
- O método `SaveToFile()` deve salvar o conteúdo do CSV em um arquivo com o nome definido.  
- Se o nome não for especificado, deve ser salvo como `"output.csv"`.  

---

## 🔹 Requisitos Não Funcionais

✅ O código deve ser limpo e organizado, seguindo boas práticas de Go.  
✅ O Builder deve permitir **encadeamento de métodos** (`SetFilename()`, `AddRow()`, etc.).  
✅ O código deve ser **extensível**, para que novos recursos possam ser adicionados no futuro.  
✅ Deve lidar com erros de forma apropriada (ex: validação do número de colunas por linha).  
