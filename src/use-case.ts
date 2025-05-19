import { inject, injectable } from 'tsyringe'
import { IRepo } from './repo'

@injectable()
export class UseCase {
  constructor(
    @inject(IRepo.name)
    private repo: IRepo,
  ) {}

  run() {
    this.repo.create()
  }
}
