import { zod } from '../utils/zod'

export const UserSchema = zod
  .object({
    id: zod.number().int().describe('ID do usuário'),
    name: zod
      .string()
      .min(3, 'Nome deve ter pelo menos 3 caracteres')
      .describe('Nome do usuário'),
    email: zod.string().email('E-mail inválido'),
  })
  .openapi('User')

export const CreateUserSchema = UserSchema.omit({ id: true })
