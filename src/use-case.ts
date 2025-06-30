import { Either } from './either'

export type UseCaseResponse<Output> = Either<Error, Output>
export abstract class UseCase<Input, Output> {
  abstract execute(input: Input): Promise<UseCaseResponse<Output>>
}

export class TestUseCase extends UseCase<number, string> {
  async execute(input: number): Promise<UseCaseResponse<string>> {
    return Either.right(String(input))
  }
}
