import express from 'express'

import { registerPostUsersRoute } from './handlers/get-users'
import { registerPutUsersRoute } from './handlers/update-user'
import { SwaggerBuilder } from './utils/swagger-builder'
import './utils/zod'
import { z } from 'zod'

const app = express()
app.use(express.json())

registerPostUsersRoute(app)
registerPutUsersRoute(app)

SwaggerBuilder.build(app)

const schema = z.array(
  z.object({
    user: z.object({
      name: z.string(),
      age: z.string(),
      isAdministrator: z.boolean(),
    }),
  }),
)

const data = [
  {
    user: {
      name: 'John Doe',
      isAdministrator: 'true',
    },
  },
  {
    user: {
      name: 1,
      isAdministrator: 'true',
    },
  },
]

const result = schema.safeParse(data)
console.error(result.error?.errors.map((e) => e.message))

const PORT = 3000
app.listen(PORT, () => {
  console.log(`Servidor rodando em http://localhost:${PORT}`)
  console.log(`Swagger disponível em http://localhost:${PORT}/api-docs`)
})
