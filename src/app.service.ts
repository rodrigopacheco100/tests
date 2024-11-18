import { Injectable, Logger } from '@nestjs/common'
import { Job, PgBossService } from '@wavezync/nestjs-pgboss'

@Injectable()
export class AppService {
  private readonly logger = new Logger(AppService.name)

  constructor(private readonly pgBoss: PgBossService) {}

  async getHello() {
    await this.pgBoss.scheduleJob('my-job', { key: 'value' })
  }

  @Job('my-job')
  async handleMyJob(job: any) {
    this.logger.log('Handling job with data:', job.data)
  }
}
