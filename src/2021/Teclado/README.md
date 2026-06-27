# Teclado

Os teclados de telefones mostram teclas com os dígitos de 0 a 9, para que possamos digitar o número do telefone que queremos contactar, como na figura abaixo. Mas as teclas também mostram letras, que podem ser usadas por exemplo para facilitar a memorização de um número de telefone em particular. Por exemplo, para memorizar o número

(74) 7622 3623

podemos associar esse número à cadeia de caracteres pipocadoce:

![](/static/img/task_images/2021f3ps_teclado.png)

Claramente, um número pode ser representado por diferentes cadeias de caracteres. Por exemplo, o número 3482 pode ser representado por fita, diva, dita, egua, e muitas outras cadeias de caracteres.

Dados um número e uma lista de cadeias de caracteres, sua tarefa é determinar quantas cadeias de caracteres da lista podem representar o número dado.

### Entrada

A primeira linha da entrada contém uma cadeia de caracteres N, o número de telefone. A segunda linha contém um inteiro M, o número de cadeias de caracteres na lista. Cada uma das M linhas seguintes contém uma cadeia de caracteres Ci.

### Saída

Seu programa deve produzir uma única linha na saída, contendo um único inteiro, o número de cadeias de caracteres da lista que podem representar o número dado.

### Restrições

-   1 ≤ comprimento de N ≤ 1 000
-   N contém apenas dígitos entre 2 e 9
-   1 ≤ M ≤ 1 000
-   Ci contém apenas letras minúsculas não acentuadas, para 1 ≤ i ≤ M
-   1 ≤ comprimento de Ci ≤ 1 000, para 1 ≤ i ≤ M
-   Ci são todas distintas para 1 ≤ i ≤ M

### Informações sobre a pontuação

-   Para um conjunto de casos de testes valendo 13 pontos, comprimento de N = 1 e M ≤ 20.
-   Para um conjunto de casos de testes valendo outros 87 pontos, nenhuma restrição adicional.

### Exemplos

**Entrada**

```bash
3482
4
fita
regua
milho
diva
```

**Saída**

```bash
2
```

---

**Entrada**

```bash
7476223623
5
pipoca
pipocadoce
misobafobe
doce
docepipoca
```

**Saída**

```bash
1
```

---

**Entrada**

```bash
4444
3
mono
tudo
nada
```

**Saída**

```bash
0
```
