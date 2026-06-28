# Mania de Ímpar

Nome do arquivo: `mania.c`, `mania.cpp`, `mania.pas`, `mania.java`, `mania.js` ou `mania.py`

Bel é uma garota muito estudiosa e inteligente. No entanto, como muitas pessoas geniais, ela tem algumas manias peculiares, sendo que a principal delas é fazer as coisas em quantidades ímpares. Isso geralmente não atrapalha a sua vida, mas às vezes cria situações interessantes. Por exemplo, Bel visitou sua tia e, como de costume, levou em sua mochila um tênis e uma meia extras, bateu na porta três vezes e tomou cinco copos de água.

Nesse dia, elas decidiram que fariam cookies. Quando a menina chegou, a massa já havia sido preparada e disposta em N linhas e M colunas em uma bandeja. A fim de atender à mania da sobrinha, a tia de Bel havia posicionado os biscoitos de forma que N e M são ímpares, mas não teve tempo de colocar as quantidades corretas de gotas de chocolate em cada cookie. Dessa forma, o biscoito na linha _i_ e coluna _j_ possui _G_\__i,j_ gotas de chocolate.

Bel decidiu modificar os cookies para estarem de acordo com sua mania. Ela considera que uma bandeja está **organizada** se, para todo par de cookies adjacentes, a soma das quantidades de gotas de chocolate nos dois cookies é ímpar. Um cookie é adjacente a outro se está imediatamente à esquerda, à direita, acima ou abaixo dele. Bel pretende adicionar gotas de chocolate em alguns cookies para deixar a bandeja organizada. Porém, para economizar os ingredientes da tia, ela deseja fazer isso adicionando o mínimo de gotas possível.

Sua tarefa é: dados os valores N e M, bem como as quantidades _G_\__i,j_ de gotas de chocolate em cada cookie, determine a quantidade mínima de gotas que precisam ser adicionadas para que a bandeja esteja organizada, isto é, para que a soma das quantidades de gotas em dois cookies adjacentes seja sempre ímpar. Além disso, você deve descrever a configuração final da bandeja organizada, isto é, indicar a quantidade de gotas de chocolate em cada cookie após as adições de Bel.

### Entrada

A primeira linha da entrada contém dois inteiros N e M, a quantidade de linhas e a quantidade de colunas na bandeja.

As próximas N linhas contêm M inteiros cada. A _i_-ésima destas linhas contém os inteiros _G_\__i,1_, _G_\__i,2_, ..., _G_\__i,M_, as quantidades de gotas de chocolate nos cookies da _i_-ésima linha.

### Saída

A primeira linha da saída deve conter um único inteiro, o mínimo de gotas de chocolate que precisam ser adicionadas para que a bandeja esteja organizada.

As N linhas seguintes devem conter M inteiros cada, indicando a configuração final da bandeja de cookies na solução ótima, no mesmo formato da entrada.

### Restrições

-   1 ≤ N, M ≤ 100
-   N e M são ímpares
-   1 ≤ _G_\__i,j_ ≤ 1.000 para todo 1 ≤ _i_ ≤ N e 1 ≤ _j_ ≤ M

### Informações sobre a pontuação

A tarefa vale 100 pontos. Estes pontos estão distribuídos em subtarefas, cada uma com suas restrições adicionais às definidas acima.

-   **Subtarefa 1 (0 pontos):** composta apenas pelos exemplos mostrados abaixo. Não vale pontos, serve apenas para verificar se o programa imprime o resultado correto para os exemplos.
-   **Subtarefa 2 (15 pontos):** N = 1 e M = 3.
-   **Subtarefa 3 (31 pontos):** N = 1.
-   **Subtarefa 4 (54 pontos):** sem restrições adicionais.

### Exemplos

**Entrada 1**

```bash
3 3
1 2 1
2 2 2
1 2 1
```

**Saída 1**

```bash
1
1 2 1
2 3 2
1 2 1
```

**Explicação:** a bandeja inicial não está organizada pois a soma dos dois primeiros cookies na segunda linha, 2 + 2, é par. Basta que Bel adicione uma gota no cookie central, de forma que ele passe a ter 3 gotas. Após a adição, a soma das gotas em dois cookies adjacentes é sempre ímpar.

---

**Entrada 2**

```bash
5 5
8 7 2 5 7
9 9 9 8 7
2 7 4 5 6
6 2 8 2 1
2 3 4 7 8
```

**Saída 2**

```bash
4
8 7 2 5 8
9 10 9 8 7
2 7 4 5 6
7 2 9 2 1
2 3 4 7 8
```

---

**Entrada 3**

```bash
1 5
1 2 3 4 5
```

**Saída 3**

```bash
0
1 2 3 4 5
```

**Explicação:** a bandeja de cookies já está organizada, logo Bel não precisa adicionar nenhuma gota de chocolate.
