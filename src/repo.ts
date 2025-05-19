export abstract class IRepo {
  abstract create: () => Promise<void>
}

export class Repository implements IRepo {
  constructor(private random = Math.random()) {}

  async create(): Promise<void> {
    console.log(this.random)
  }
}
