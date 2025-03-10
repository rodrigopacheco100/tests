import { Schema } from 'swagger-jsdoc'

export const HandlerOutput: Schema = {
  type: 'object',
  properties: {
    message: { type: 'string' },
  },
}
