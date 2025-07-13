import { Effect, Data, pipe, Resource } from 'effect'

{
  // Criando um Effect Simples

  // Effect básico que sempre tem sucesso
  const helloWorld = Effect.succeed('Hello, World!')

  // Effect que pode falhar
  const divide = (a: number, b: number) =>
    b === 0 ? Effect.fail(new Error('Division by zero')) : Effect.succeed(a / b)

  // Executando Effects
  Effect.runPromise(helloWorld).then(console.log) // "Hello, World!"
  Effect.runPromise(divide(10, 2)).then(console.log) // 5
  // Effect.runPromise(divide(10, 0)).then(console.log) // Error example
}

{
  // Tratamento de Erros Estruturado

  // Definindo tipos de erro customizados
  class ValidationError extends Data.TaggedError('ValidationError')<{
    message: string
  }> {}

  class NetworkError extends Data.TaggedError('NetworkError')<{
    status: number
    message: string
  }> {}

  // Função que pode falhar com diferentes tipos de erro
  const validateEmail = (
    email: string,
  ): Effect.Effect<string, ValidationError | NetworkError> => {
    if (!email.includes('@')) {
      return Effect.fail(
        new ValidationError({ message: 'Invalid email format' }),
      )
    }
    return Effect.succeed(email)
  }

  // Tratando erros específicos
  const program = validateEmail('invalid-email').pipe(
    Effect.catchTag('ValidationError', (error) =>
      Effect.succeed(`Fixed: ${error.message}`),
    ),
  )

  Effect.runPromise(program).then(console.log) // "Fixed: Invalid email format"
}

{
  // Composição de Effects

  // Simulando operações assíncronas
  const fetchUser = (id: number) =>
    Effect.succeed({ id, name: 'John Doe', email: 'john@example.com' })

  const fetchUserPosts = (userId: number) =>
    Effect.succeed([
      { id: 1, title: 'Post 1', content: 'Content 1' },
      { id: 2, title: 'Post 2', content: 'Content 2' },
    ])

  // Compondo Effects sequencialmente
  const getUserWithPosts = (userId: number) =>
    pipe(
      fetchUser(userId),
      Effect.flatMap((user) =>
        pipe(
          fetchUserPosts(user.id),
          Effect.map((posts) => ({ user, posts })),
        ),
      ),
    )

  // Executando a composição
  Effect.runPromise(getUserWithPosts(2)).then(console.log)
}
{
  // Simulando uma conexão de banco de dados
  class DatabaseConnection {
    constructor(private url: string) {}

    query(sql: string) {
      return `Result for: ${sql}`
    }

    close() {
      console.log('Database connection closed')
    }
  }

  // Criando um recurso gerenciado
  const acquireDatabase = (url: string) =>
    Effect.acquireRelease(
      Effect.sync(() => {
        console.log('Opening database connection')
        return new DatabaseConnection(url)
      }),
      (db) => Effect.sync(() => db.close()),
    )

  // Usando o recurso
  const program = Effect.scoped(
    acquireDatabase('postgresql://localhost').pipe(
      Effect.flatMap((db) =>
        Effect.sync(() => db.query('SELECT * FROM users')),
      ),
    ),
  )

  Effect.runPromise(program).then(console.log)
}
