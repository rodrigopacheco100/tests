import { DateTime } from 'luxon'

const now = new Date()
console.log('🚀 ~ now:', now)

const classStartsInDateTime =
  DateTime.fromJSDate(now).setZone('America/Sao_Paulo')

const a = classStartsInDateTime.minus({ day: 1 }).startOf('day')
const b = classStartsInDateTime.plus({ day: 1 }).endOf('day')

console.log(a)
console.log(b)
console.log(a.diff(b))
console.log(a < b)
