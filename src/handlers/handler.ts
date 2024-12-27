export type HandlerInput = {
  message: string
}

export type HandlerOutput = {
  message: string
}

/**
 * @openapi
 * /message/{message}:
 *   get:
 *     summary: Retorna a mensagem fornecida.
 *     description: Recebe uma mensagem como parâmetro e a retorna.
 *     parameters:
 *       - name: message
 *         in: path
 *         required: true
 *         description: Mensagem a ser retornada.
 *         schema:
 *           type: string
 *       - name: message2
 *         in: query
 *         required: true
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Mensagem retornada com sucesso.
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 message:
 *                   type: string
 *                   description: Mensagem solicitada.
 */
export class Handler {
  public perform(input: HandlerInput): HandlerOutput {
    return input
  }
}
