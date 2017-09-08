const gql = require('graphql-tag')

module.exports = gql`
  subscription {
    Trigger(
      filter: {
        mutation_in: [CREATED, UPDATED]
      }
    ) {
      node {
        contract {
          abi
          address
        }
        id
        enable
        eventName
      }
    }
  }`