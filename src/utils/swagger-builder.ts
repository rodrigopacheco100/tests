import {
  OpenApiGeneratorV3,
  OpenAPIRegistry as ZodOpenApiRegistry,
  RouteConfig,
} from '@asteasolutions/zod-to-openapi'
import { Application as ExpressApplication } from 'express'
import { serve, setup } from 'swagger-ui-express'

export type PathConfig = RouteConfig

export abstract class SwaggerBuilder {
  private static registry = new ZodOpenApiRegistry()

  public static registerPath(options: PathConfig) {
    this.registry.registerPath(options)
  }

  public static build(app: ExpressApplication) {
    const generator = new OpenApiGeneratorV3(this.registry.definitions)
    const openapiDocs = generator.generateDocument({
      openapi: '3.0.0',
      info: {
        title: 'API de Usuários',
        version: '1.0.0',
      },
    })

    app.use('/api-docs', serve, setup(openapiDocs))
  }
}
