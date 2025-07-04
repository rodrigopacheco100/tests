import PDFDocument from 'pdfkit'
import fs from 'fs'
import { performance } from 'perf_hooks'

type BenchmarkResult = {
  rows: number
  timeMs: number
  timeSec: number
  ramUsedMB: number
  ramPeakMB: number
  cpuUserMs: number
  cpuSystemMs: number
  fileSizeMB: number
  throughput: number
}

// Função para obter uso de memória
function getMemoryUsage() {
  const used = process.memoryUsage()
  return {
    rss: Math.round((used.rss / 1024 / 1024) * 100) / 100, // MB
    heapUsed: Math.round((used.heapUsed / 1024 / 1024) * 100) / 100, // MB
    heapTotal: Math.round((used.heapTotal / 1024 / 1024) * 100) / 100, // MB
    external: Math.round((used.external / 1024 / 1024) * 100) / 100, // MB
  }
}

// Função para executar benchmark
async function runBenchmark(rows: number): Promise<BenchmarkResult> {
  return new Promise(async (resolve) => {
    console.log(`\n📊 Testando com ${rows.toLocaleString()} linhas...`)

    // Força garbage collection se disponível
    if (global.gc) global.gc()

    const startMemory = getMemoryUsage()
    const startTime = performance.now()
    const startCPU = process.cpuUsage()

    const doc = new PDFDocument()
    const filename = `benchmark-${rows}.pdf`
    doc.pipe(fs.createWriteStream(filename))

    const table = doc.table({
      columnStyles: {
        backgroundColor: '#f0f0f0',
      },
    })

    // Adiciona linhas em lotes para melhor performance
    const batchSize = 1000
    for (let i = 0; i < rows; i += batchSize) {
      const currentBatch = Math.min(batchSize, rows - i)
      for (let j = 0; j < currentBatch; j++) {
        table.row([
          `Linha ${i + j + 1}`,
          `Dados B${i + j + 1}`,
          `Dados C${i + j + 1}`,
        ])
      }

      // Permite que o event loop processe outras tarefas
      if (i % 5000 === 0) {
        await new Promise((resolve) => setImmediate(resolve))
      }
    }

    doc.end()

    doc.on('end', () => {
      const endTime = performance.now()
      const endCPU = process.cpuUsage(startCPU)
      const endMemory = getMemoryUsage()

      // Obtém tamanho do arquivo
      const stats = fs.statSync(filename)
      const fileSizeMB = Math.round((stats.size / 1024 / 1024) * 100) / 100

      const result = {
        rows,
        timeMs: Math.round((endTime - startTime) * 100) / 100,
        timeSec: Math.round(((endTime - startTime) / 1000) * 100) / 100,
        ramUsedMB: endMemory.heapUsed - startMemory.heapUsed,
        ramPeakMB: endMemory.heapUsed,
        cpuUserMs: Math.round(endCPU.user / 1000),
        cpuSystemMs: Math.round(endCPU.system / 1000),
        fileSizeMB,
        throughput: Math.round(rows / ((endTime - startTime) / 1000)),
      }

      // Remove arquivo após medição
      // fs.unlinkSync(filename)

      resolve(result)
    })
  })
}

// Configurações de teste
const testCases = [
  1000, // 1K linhas
  5000, // 5K linhas
  10000, // 10K linhas
  25000, // 25K linhas
  50000, // 50K linhas
  100000, // 100K linhas
]

console.log('🚀 Iniciando benchmark PDFKit')
console.log('═'.repeat(80))
//
;(async () => {
  const results: BenchmarkResult[] = []
  // Executa testes sequencialmente
  for (const testCase of testCases) {
    try {
      const result = await runBenchmark(testCase)
      results.push(result)

      console.log(
        `✅ ${testCase.toLocaleString()} linhas: ${result.timeSec}s | RAM: ${result.ramPeakMB}MB | Arquivo: ${result.fileSizeMB}MB`,
      )
    } catch (error) {
      console.error(`❌ Erro com ${testCase} linhas:`, (error as Error).message)
    }
  }

  // Exibe tabela de resultados
  console.log('\n📈 RESULTADOS DO BENCHMARK')
  console.log('═'.repeat(120))
  console.log(
    '┌─────────────┬──────────────┬────────────┬─────────────┬─────────────┬──────────────┬───────────────┐',
  )
  console.log(
    '│    Linhas   │    Tempo     │    RAM     │   Arquivo   │  CPU User   │ CPU System   │  Throughput   │',
  )
  console.log(
    '│             │     (s)      │    (MB)    │    (MB)     │    (ms)     │    (ms)      │  (linhas/s)   │',
  )
  console.log(
    '├─────────────┼──────────────┼────────────┼─────────────┼─────────────┼──────────────┼───────────────┤',
  )

  results.forEach((r) => {
    const linha = `│ ${r.rows.toLocaleString().padStart(11)} │ ${r.timeSec.toString().padStart(12)} │ ${r.ramPeakMB.toString().padStart(10)} │ ${r.fileSizeMB.toString().padStart(11)} │ ${r.cpuUserMs.toString().padStart(11)} │ ${r.cpuSystemMs.toString().padStart(12)} │ ${r.throughput.toLocaleString().padStart(13)} │`
    console.log(linha)
  })

  console.log(
    '└─────────────┴──────────────┴────────────┴─────────────┴─────────────┴──────────────┴───────────────┘',
  )

  // Análise de escalabilidade
  if (results.length > 1) {
    console.log('\n📊 ANÁLISE DE ESCALABILIDADE')
    console.log('═'.repeat(50))

    const firstResult = results[0]
    const lastResult = results[results.length - 1]

    const scaleFactorRows = lastResult.rows / firstResult.rows
    const scaleFactorTime = lastResult.timeSec / firstResult.timeSec
    const scaleFactorRAM = lastResult.ramPeakMB / firstResult.ramPeakMB

    console.log(`📈 Fator de escala (linhas): ${scaleFactorRows}x`)
    console.log(
      `⏱️  Fator de escala (tempo): ${Math.round(scaleFactorTime * 100) / 100}x`,
    )
    console.log(
      `💾 Fator de escala (RAM): ${Math.round(scaleFactorRAM * 100) / 100}x`,
    )

    if (scaleFactorTime < scaleFactorRows) {
      console.log(
        '✅ Performance escalável: tempo cresce menos que proporcionalmente',
      )
    } else {
      console.log(
        '⚠️  Performance não-linear: tempo cresce mais que proporcionalmente',
      )
    }
  }

  console.log('\n🎯 Benchmark concluído!')
})()
