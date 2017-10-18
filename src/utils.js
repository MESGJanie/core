const Logger = require('./logger')
const { ConnectionError } = require('./errors')

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms))

const testConnection = async (check, endpoint) => {
  let retry = process.env.RETRY_COUNT || 10
  const delay = process.env.RETRY_DELAY || 1000
  const time = new Date()
  while (!check()) {
    if (retry <= 0) {
      throw new ConnectionError(endpoint)
    }
    await sleep(delay)
    Logger.warn(`Connection to ${endpoint} invalid... retrying in ${delay}ms`)
    retry = retry - 1
  }
  return Promise.resolve({
    retryCount: retry - process.env.RETRY_COUNT,
    ms: +new Date() - time
  })
}

module.exports = {
  sleep,
  testConnection
}
