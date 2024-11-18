import { Module } from '@nestjs/common'
import { AppController } from './app.controller'
import { AppService } from './app.service'
import { PgBossModule } from '@wavezync/nestjs-pgboss'

@Module({
  imports: [
    PgBossModule.forRootAsync({
      useFactory: async () => ({
        connectionString: 'postgresql://docker:docker@localhost:5432/tests',
      }),
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
