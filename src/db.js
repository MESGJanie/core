const ws = require('ws')
const { ApolloClient, createNetworkInterface } = require('apollo-client')
const { SubscriptionClient, addGraphQLSubscriptions } = require('subscriptions-transport-ws')
const queryFetchAll = require('./queries/fetchAll')
const queryCreated = require('./queries/created')
const queryUpdated = require('./queries/updated')
const queryDeleted = require('./queries/deleted')
const mutationEvent = require('./queries/mutationEvent')

const headers = {
  Authorization: [
    'Bearer',
    process.env.AUTH_TOKEN
  ].join(' ')
}

const client = new ApolloClient({
  networkInterface: addGraphQLSubscriptions(
    createNetworkInterface({
      uri: process.env.GRAPHQL_HTTP_ENDPOINT,
      opts: {
        headers,
      }
    }),
    new SubscriptionClient(process.env.GRAPHQL_WS_ENDPOINT, {
      reconnect: true,
      connectionParams: headers,
    }, ws)
  )
})

const fetchAll = () => client
  .query({ query: queryFetchAll })
  .then(x => x.data.allTriggers)

const onDataCreated = callback => client
  .subscribe({ query: queryCreated })
  .subscribe({
    next: value => callback(null, value.Trigger.node),
    error: error => callback(error, null)
  })

const onDataUpdated = callback => client
  .subscribe({ query: queryUpdated })
  .subscribe({
    next: value => callback(null, value.Trigger.node),
    error: error => callback(error, null)
  })

const onDataDeleted = callback => client
  .subscribe({ query: queryDeleted })
  .subscribe({
    next: value => callback(null, value.Trigger.previousValues.id),
    error: error => callback(error, null)
  })

const writeEvent = (event, trigger) => client
  .mutate({
    mutation: mutationEvent,
    variables: {
      payload: event.args,
      transactionId: event.transactionHash,
      triggerId: trigger.id
    }
  })
  .then(x => x.data.createEvent)

module.exports = {
  fetchAll,
  onDataCreated,
  onDataUpdated,
  onDataDeleted,
  writeEvent
}