import { extendZodWithOpenApi } from '@asteasolutions/zod-to-openapi'
import { z } from 'zod'
import { buildErrorMap } from 'zod-error-utils'

extendZodWithOpenApi(z)
const customErrorMap = buildErrorMap({
  prefixFn(path, message) {
    return `${path}: ${message}`
  },
})

z.setErrorMap(customErrorMap)
export { z as zod }
