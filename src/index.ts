import os from 'node:os'
import express from 'express'

const app = express()
app.use(express.json())

app.get('/', (req, res) => {
  const totalMemory = (os.totalmem() / 1024 / 1024 / 1024)
    .toFixed(2)
    .concat(' GB')
  const freeMemory = (os.freemem() / 1024 / 1024 / 1024)
    .toFixed(2)
    .concat(' GB')
  const CPUCores = os.cpus().length

  return res.json({
    operatingSystem: os.type(),
    architecture: os.arch(),
    platform: os.platform(),
    release: os.release(),
    freeMemory,
    totalMemory,
    CPUCores,
  })
})

app.listen(3333, () => {
  console.log('🚀 Express server running!')
})
