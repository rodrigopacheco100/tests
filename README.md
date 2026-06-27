<p align="center">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.svg" alt="Gopher" width="100"/>
    <br/>
    <img src="https://img.shields.io/badge/OBI-Olimp%C3%ADada%20Brasileira%20de%20Inform%C3%A1tica-009739?style=for-the-badge&labelColor=002776" alt="OBI"/>
</p>

<h1 align="center">
    🏆 OBI Solutions · Golang
</h1>

<p align="center">
    <b>Soluções de exercícios e provas da</b><br>
    <b>Olimpíada Brasileira de Informática</b><br>
    escritas em <b>Go</b>
</p>

<p align="center">
    <a href="#-sobre">Sobre</a> •
    <a href="#-estrutura">Estrutura</a> •
    <a href="#-como-executar">Como executar</a>
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Linguagem-Go-00ADD8?style=flat-square&logo=go" alt="Go"/>
    <img src="https://img.shields.io/badge/Status-Em%20desenvolvimento-yellow?style=flat-square" alt="Status"/>
    <img src="https://img.shields.io/badge/Fase-Única%20(simulados)-purple?style=flat-square" alt="Fase"/>
    <img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="MIT"/>
</p>

<br />

## 📚 Sobre

Este repositório reúne minhas soluções para problemas da **[Olimpíada Brasileira de Informática (OBI)](https://olimpiada.ic.unicamp.br/)**, resolvidos utilizando **Go (Golang)**.

A OBI é uma competição organizada pelo **Instituto de Computação da UNICAMP** que estimula o raciocínio lógico-matemático e a criatividade na solução de problemas computacionais entre estudantes brasileiros.

> 🎯 **Objetivo:** documentar minha jornada de estudos, compartilhar abordagens e manter um histórico de soluções que possam ajudar outros competidores.

<br />

## 🗂️ Estrutura

```
📦 src/
 ┗━ 📂 <ano>/
      ┗━ 📂 <nome do problema>/
           ┗━ 📜 main.go
```

Os problemas são organizados por **ano** da prova. Cada problema contém seu próprio diretório com o código-fonte em Go.

```
src/
└── 2021/
      └── Tempo de resposta/
            └── main.go
```

<br />

## ▶️ Como Executar

Certifique-se de ter o [Go](https://go.dev/dl/) instalado (versão 1.21+):

```bash
# Executar um problema específico
go run src/2021/Tempo-de-resposta/main.go

# Ou usando entrada por arquivo
go run src/2021/Tempo-de-resposta/main.go < src/2021/Tempo-de-resposta/examples/example1.txt
```

## 🧪 Testes

Alguns problemas podem incluir arquivos de teste para verificar as soluções:

```bash
make test
```

<br />

<p align="center">
    Feito com ❤️ e ☕ — e muito Go!<br>
</p>
