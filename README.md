# Tutorial para testar o Github Actions

1. Executar o comando abaixo para iniciar um projeto em GO.

```SHELL
$ go mod init github.com/...
```

2. Criar o arquivo `main.go` na raiz do projeto.

3. Inserir o seguinte código dentro da `main.go`:

```GO
package main

import "fmt"

func main() {
	fmt.Println("Resultado da soma de 1 + 2:", sum(1, 2))
}

// O código está propositalmente incorreto para exemplificar uma pipeline que rodará com erros.
func sum(a, b int) int {
	return a - b
}
```

4. Criar o arquivo `main_test.go` também na raiz do projeto.

5. Inserir o seguinte código dentro da `main_test.go`:

```GO
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Teste básico para verificar se a soma de 1 + 2 é igual a 3. Ao executar o deste, desta vez, ele deve falhar.
func TestSum(t *testing.T) {
	assert.Equal(t, 3, sum(1, 2), "1 + 2 should be 3")
}
```

6. Executar o comando abaixo para baixar/atualizar as dependências:

```SHELL
$ go mod tidy
```

7. Criar o diretório `.github/workflows` com o arquivo `ci.yaml`.

8. Dentro do arquivo `ci.yaml` colar o seguinte conteúdo:

```YAML
name: Continuous Integration

on:
  push: # Quando acontecer um push
    branches: # na Branch
      - master # master
  # ou
  pull_request: # Quando acontecer um pull request
    branches: # na branch
      - master # master

jobs: # Definição de Jobs
  continuous-integration: # Nome do Job
    runs-on: ubuntu-latest # Runner

    steps:
      - name: Checkout code # Nome do passo
        uses: actions/checkout@v2 # Utiliza uma action do marketplace para fazer checkout do código

      - name: Setup Go # Nome do passo
        uses: actions/setup-go@v5 # Utiliza uma action do marketplace para configurar o ambiente go
        with: # Configuração da action
          go-version: '1.23.2' # Versão do Go

      - name: Install dependencies # Nome do passo
        run: go mod download # Executa o comando go mod download para baixar as dependências

      - name: Run tests # Nome do passo
        run: go test ./... -v # Executa o comando go test ./... -v para rodar os testes

      - name: Build # Nome do passo
        run: go build ./... # Executa o comando go build ./... para compilar o código
```

9. Subir o código no github e verificar o status da pipeline.

10. Criar uma branch a partir da master da seguinte forma:

```SHELL
$ git checkout -b branch-from-master
```

11. Corrigir a função `sum()` do arquivo `main.go`:

```GO
package main

import "fmt"

func main() {
	fmt.Println("Resultado da soma de 1 + 2:", sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}
```

12. Commitar e fazer push para o repositório remoto.

13. Abrir um pull request apontando para o master. (Neste passo a pipeline deve rodar).

14. Verificar status ad pipeline.

15. Fazer merge do PR no master. (Neste passo a pipeline deve rodar novamente).