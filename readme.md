# bear
**bear** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

### Commands to set-up chain
```
ignite scaffold chain bear

ignite scaffold type bear role background clothes weapon creator id:uint

ignite scaffold message create-bear role background clothes weapon --response id:uint

ignite scaffold message update-bear role background clothes weapon id:uint

ignite scaffold message delete-bear id:uint

ignite scaffold query show-bear id:uint --response bear:Bear

ignite scaffold query list-bear --response bear:Bear --paginated

ignite scaffold query list-bear-role role --response bear:Bear --paginated

ignite scaffold query list-bear-background background --response bear:Bear --paginated

ignite scaffold query list-bear-clothes clothes --response bear:Bear --paginated

ignite scaffold query list-bear-weapon weapon --response bear:Bear --paginated

ignite scaffold query list-bear-creator creator --response bear:Bear --paginated
```

### Commands for queries
```
# Create bears
beard tx bear create-bear warrior red armour sword --from alice --chain-id bear
beard tx bear create-bear healer blue t-shirt wand --from bob --chain-id bear

# Show a specific bear based on id
beard q bear show-bear 0
beard q bear show-bear 1

# List all bears
beard q bear list-bear

# List all bears with a specific role
beard q bear list-bear-role warrior
beard q bear list-bear-role healer

# List all bears with specific background
beard q bear list-bear-background red
beard q bear list-bear-background blue

# List all bears with specific clothes
beard q bear list-bear-clothes armour
beard q bear list-bear-clothes t-shirt

# List all bears with specific weapon
beard q bear list-bear-weapon sword
beard q bear list-bear-weapon wand

# List all bears that are created by a specified address
beard q bear list-bear-creator {alice_address}
beard q bear list-bear-creator {bob_address}

# Update the attributes of a specific bear
beard tx bear update-bear mage green cloak stick 0 --from alice --chain-id bear

# Delete a specific bear
beard tx bear delete-bear 0 --from alice --chain-id bear
beard tx bear delete-bear 1 --from bob --chain-id bear
```

## Consensus Breaking Change
- Breaking consensus refers to causing the different nodes in the blockchain network to not be able to agree on the actual state of the blockchain, even though the transaction is valid. This could be due to the different nodes in the blockchain network having a different state.

- In order to break the consensus, I have chosen to change the way in which the node processes create-bear and delete-bear messages.

- On nodes with the breaking change, a bear with an empty role, background, clothes and weapon will be created when a create-bear transaction message is received even though the role, background, clothes and weapon is provided.

- Similarly, on nodes with the breaking change, the bear specified in the delete-bear transaction will not be deleted.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/bear@latest! | sudo bash
```
`username/bear` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
