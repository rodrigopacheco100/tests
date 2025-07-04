import PDFDocument from 'pdfkit'
import fs from 'fs'

import { criarGraficoBarras } from './helpers/criar-grafico-barra'

// Simula dados vindos de uma API ou banco de dados
async function gerarRelatorioDinamico() {
  const doc = new PDFDocument({ size: 'A4' })
  doc.pipe(fs.createWriteStream('relatorio-dinamico.pdf'))

  // Dados dinâmicos (pode vir de qualquer fonte)
  const vendas = [
    { mes: 'Jan', vendas: 1200, gastos: 800 },
    { mes: 'Feb', vendas: 1900, gastos: 1200 },
    { mes: 'Mar', vendas: 50000, gastos: 1800 },
    { mes: 'Apr', vendas: 1500, gastos: 1000 },
    { mes: 'May', vendas: 2500, gastos: 1500 },
  ]

  doc.fontSize(20).text('Relatório de Vendas', { align: 'center' })
  doc.moveDown()

  // Gráfico de barras para vendas
  const dadosVendas = vendas.map((v) => ({
    label: v.mes,
    valor: v.vendas,
    cor: '#4CAF50',
  }))

  criarGraficoBarras(doc, dadosVendas, 0, 100, 600, 200)

  doc.end()
}

gerarRelatorioDinamico()
