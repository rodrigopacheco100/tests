import { Either } from './either'

class ForbiddenError extends Error {
  constructor() {
    super('Forbidden')
  }
}

function main(message: string): Either<Error, string> {
  if (message !== 'hello') return Either.left(new ForbiddenError())

  return Either.right('world')
}

const result = main('batata')

if (result.isLeft()) throw result.value

console.log(result.value)
