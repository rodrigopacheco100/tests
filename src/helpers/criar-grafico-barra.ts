export function criarGraficoBarras(
  doc: PDFKit.PDFDocument,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  dados: any[],
  x: number,
  y: number,
  largura: number,
  altura: number,
) {
  const margem = 10
  const larguraUtil = largura - margem * 2
  const alturaUtil = altura - margem * 2

  // Encontra o valor máximo para escala
  const maxValor = Math.max(...dados.map((d) => d.valor))

  // Desenha eixos
  doc
    .strokeColor('#000000')
    .lineWidth(1)
    .moveTo(x + margem, y + margem)
    .lineTo(x + margem, y + altura - margem) // Eixo Y
    .lineTo(x + largura - margem, y + altura - margem) // Eixo X
    .stroke()

  // Desenha barras
  const larguraBarra = (larguraUtil / dados.length) * 0.8
  const espacamento = larguraUtil / dados.length

  dados.forEach((item, index) => {
    const alturaBarra = (item.valor / maxValor) * alturaUtil
    const posX =
      x + margem + index * espacamento + (espacamento - larguraBarra) / 2
    const posY = y + altura - margem - alturaBarra

    // Desenha a barra
    doc
      .fillColor(item.cor || '#4CAF50')
      .rect(posX, posY, larguraBarra, alturaBarra)
      .fill()

    // Adiciona rótulo
    doc
      .fillColor('#000000')
      .fontSize(8)
      .text(item.label, posX, y + altura - margem + 5, {
        width: larguraBarra,
        align: 'center',
      })

    // Adiciona valor
    doc.text(item.valor.toString(), posX, posY - 15, {
      width: larguraBarra,
      align: 'center',
    })
  })
}
