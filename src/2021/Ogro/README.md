# Ogro

O Ogro da Nlogônia está aprendendo a contar até dez usando os dedos das mãos (assim como os humanos, ele possui 2 mãos com 5 dedos cada). Ele está treinando muito, mas gostaria de ter um aplicativo para ajudá-lo nessa empreitada. O Ogro aprendeu a mostrar a representação de um número com as mãos da seguinte forma:

-   se o número pode ser representado usando apenas uma das mãos, o Ogro usa os dedos na mão esquerda e mantém a mão direita fechada.
-   caso contrário, o Ogro mostra todos os cinco dedos da mão esquerda, e na mão direita mostra os dedos que faltam para representar o número.

Por exemplo, para o número 3, o Ogro mostra:

```III```

onde cada letra ```I``` representa um dedo e a mão fechada é representada pelo símbolo ```*``` (asterisco). Para o número 8 o Ogro mostra:

```IIIII III```

Sua tarefa é ajudar o Ogro em seu treinamento, escrevendo um programa para, dado um número entre 0 e 10, mostrar a configuração de dedos correspondente a esse número, de acordo com as regras acima.

### Entrada

A primeira e única linha da entrada contém um inteiro N, o número que deve ser representado com os dedos das mãos.

### Saída

Seu programa deve produzir duas linhas na saída. A primeira linha deve conter a representação dos dedos da mão esquerda, a segunda linha deve conter a representação dos dedos da mão direita. A letra ```I``` deve ser usada para representar um dedo, e o caractere ```*``` (asterisco) deve ser usado para representar a mão fechada (isto é, nenhum dedo mostrado).

### Restrições

-   0 ≤ N ≤ 10

### Exemplos

**Entrada**

```bash
8
```

**Saída**

```bash
IIIII
III
```


**Entrada**

```bash
3
```

**Saída**

```bash
III
*
```


**Entrada**

```bash
0
```

**Saída**

```bash
*
*
```
