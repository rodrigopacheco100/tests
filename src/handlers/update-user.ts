import { CreateUserSchema, UserSchema } from '../schemas/user'
import { Application } from 'express'
import { SwaggerBuilder } from '../utils/swagger-builder'
import { zod } from '../utils/zod'

SwaggerBuilder.registerPath({
  method: 'put',
  path: '/users/{id}',
  tags: ['Usuários'],
  request: {
    params: zod.object({
      id: zod.number().int(),
    }),
    body: {
      content: {
        'application/json': {
          schema: CreateUserSchema,
        },
      },
    },
  },
  responses: {
    200: {
      description: 'Lista de usuários',
      content: {
        'application/json': {
          schema: zod.array(UserSchema),
        },
      },
    },
  },
})

const users: unknown[] = []

export function registerPutUsersRoute(app: Application) {
  app.put('/users/:id', (req, res) => {
    res.json(users)
  })
}
