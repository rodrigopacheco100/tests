export function criarGraficoLinha(
  doc: PDFKit.PDFDocument,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  dados: any[],
  x: number,
  y: number,
  largura: number,
  altura: number,
) {
  const margem = 40
  const larguraUtil = largura - margem * 2
  const alturaUtil = altura - margem * 2

  const maxValor = Math.max(...dados.map((d) => d.valor))
  const minValor = Math.min(...dados.map((d) => d.valor))
  const intervalo = maxValor - minValor

  // Desenha eixos
  doc
    .strokeColor('#CCCCCC')
    .lineWidth(1)
    .moveTo(x + margem, y + margem)
    .lineTo(x + margem, y + altura - margem)
    .lineTo(x + largura - margem, y + altura - margem)
    .stroke()

  // Desenha grid
  for (let i = 1; i <= 5; i++) {
    const gridY = y + margem + (alturaUtil / 5) * i
    doc
      .strokeColor('#EEEEEE')
      .moveTo(x + margem, gridY)
      .lineTo(x + largura - margem, gridY)
      .stroke()
  }

  // Desenha linha
  doc.strokeColor('#2196F3').lineWidth(2)

  dados.forEach((ponto, index) => {
    const pontoX = x + margem + (index / (dados.length - 1)) * larguraUtil
    const pontoY =
      y + altura - margem - ((ponto.valor - minValor) / intervalo) * alturaUtil

    if (index === 0) {
      doc.moveTo(pontoX, pontoY)
    } else {
      doc.lineTo(pontoX, pontoY)
    }

    // Adiciona ponto
    doc.save().fillColor('#2196F3').circle(pontoX, pontoY, 3).fill().restore()

    // Adiciona rótulo
    doc
      .fillColor('#000000')
      .fontSize(8)
      .text(ponto.label, pontoX - 10, y + altura - margem + 5)
  })

  doc.stroke()
}
