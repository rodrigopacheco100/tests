# Arara

Nome do arquivo: `arara.c`, `arara.cpp`, `arara.pas`, `arara.java`, `arara.js` ou `arara.py`

Consideradas por muitos como símbolos da fauna brasileira, as araras são aves que se destacam pelas suas penas coloridas, animação e inteligência. Assim, a chegada de novas araras ao zoológico de São Paulo é sempre muito aguardada pelo público.

Esta semana, o zoológico deseja fazer um evento para apresentar suas N araras ao público. Na região do zoológico onde o evento vai acontecer, existem M gaiolas alinhadas, numeradas de 1 a M da esquerda para a direita. Cada gaiola pode abrigar uma única arara ou ficar vazia. Os funcionários gostariam de distribuir as N araras entre as M gaiolas para o evento.

No entanto, as araras têm a tendência de expressar emoções gritando _Arara!_ de maneira inesperada. Quando uma arara grita, outras araras que estejam em gaiolas muito próximas podem se assustar com o barulho e também começar a gritar, o que por sua vez pode assustar outras araras e assim por diante.

Felizmente, pesquisadores do zoológico descobriram que uma arara se assusta com o grito de outra arara somente se existem menos do que quatro gaiolas entre elas (desconsiderando suas próprias gaiolas). Por exemplo, o grito de uma arara na gaiola 8 assustaria uma arara na gaiola 6 (existe apenas uma gaiola entre elas) ou uma arara na gaiola 12 (existem apenas três gaiolas entre elas), mas não assustaria uma arara na gaiola 3 (existem quatro gaiolas entre elas) ou na gaiola 16 (existem sete gaiolas entre elas).

O zoológico decidiu que o evento pode ser realizado se é possível distribuir as N araras entre as M gaiolas de modo que nenhum grito _Arara!_ possa assustar outras araras, ou seja, de modo que existam pelo menos quatro gaiolas vazias entre quaisquer duas araras. Sua tarefa é determinar se o evento pode ser realizado.

### Entrada

A entrada possui uma única linha contendo dois inteiros N e M indicando, respectivamente, o número de araras e o número de gaiolas.

### Saída

Seu programa deverá imprimir uma única linha contendo um único caractere. Caso seja possível distribuir as araras entre as gaiolas para realizar o evento, imprima o caractere `S` (a letra S maiúscula). Caso seja impossível, imprima o caractere `N` (a letra N maiúscula).

### Restrições

-   1 ≤ N ≤ 1.000
-   1 ≤ M ≤ 1.000

### Informações sobre a pontuação

-   A tarefa vale 100 pontos.

### Exemplos

**Entrada 1**

```bash
2 10
```

**Saída 1**

```bash
S
```

**Explicação:** é possível distribuir as duas araras da seguinte forma: uma na gaiola 1 e outra na gaiola 10. Como há 8 gaiolas vazias entre elas, caso uma arara grite, não irá assustar a outra.

---

**Entrada 2**

```bash
3 10
```

**Saída 2**

```bash
N
```

---

**Entrada 3**

```bash
3 11
```

**Saída 3**

```bash
S
```

**Explicação:** é possível distribuir as três araras da seguinte forma: uma na gaiola 1, outra na gaiola 6 e a última na gaiola 11. Perceba que, entre quaisquer duas araras, há pelo menos 4 gaiolas vazias entre elas.
