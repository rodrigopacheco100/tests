import { CreateUserSchema, UserSchema } from '../schemas/user'
import { Application } from 'express'
import { SwaggerBuilder } from '../utils/swagger-builder'
import { zod } from '../utils/zod'

SwaggerBuilder.registerPath({
  method: 'post',
  path: '/users',
  tags: ['Usuários'],
  request: {
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

export function registerPostUsersRoute(app: Application) {
  app.post('/users', (req, res) => {
    res.json(users)
  })
}
