import express from 'express'
import swaggerJSDoc from 'swagger-jsdoc'
import swaggerUi from 'swagger-ui-express'
import { Handler } from './handlers/handler'

const app = express()
app.use(express.json())

const options: swaggerJSDoc.Options = {
  definition: {
    info: {
      title: 'Exemplo API',
      version: '1.0.0',
      description: 'API de exemplo usando Swagger JSDoc',
    },
  },
  apis: ['./src/**/*.ts'],
}

const swaggerSpec = swaggerJSDoc(options)

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec))

app.get('/message/:message', (req, res) => {
  const handler = new Handler()

  const result = handler.perform({ message: req.params.message })

  return res.json(result)
})

app.listen(3333, () => {
  console.log('🚀 Express server running!')
})
