# Cosmos Indexer Modules

The Cosmos Indexer Modules is a set of open-source Go packages designed to extend the [Cosmos Indexer](https://github.com/DefiantLabs/cosmos-indexer) application. They provide blockchain/custom module specific indexing capabilities to the Cosmos Indexer application. This allows the Cosmos Indexer to query, parse and index custom modules on the Cosmos SDK-based blockchains beyond the core SDK modules and their built in message types.

## The Problem Overview

### Cosmos SDK, Custom Blockchain Modules, and Custom Transaction Message Types

The Cosmos SDK is a powerful framework for building blockchain applications. The SDK allows developers to create custom modules that can be added to the blockchain to extend its functionality. These custom modules can have their own Transaction message types.

This provides a powerful way to extend the core functionality of the blockchain, but it also presents a challenge when it comes to indexing and querying the blockchain data. Each blockchain can have its own set of custom modules with their own message types and events, and the Cosmos Indexer needs to be able to parse and index these custom modules in a generalized way.

### Cosmos Indexer Limitations

The Cosmos Indexer is a tool that indexes the blockchain data to a generalized transaction/event database schema. However, in its default configuration, the Cosmos Indexer only indexes the core SDK modules and their built in message types. It only knows how to parse and index the core SDK messages and events provided by the modules that come with the SDK, due to the way the Cosmos Indexer type codec handles messages by type URL (see the Cosmos Indexer [Probe Walkthrough](https://github.com/DefiantLabs/cosmos-indexer/blob/main/docs/reference/custom_cosmos_module_extensions/probe_codec_walkthrough.md) for details).

This means that if a blockchain has custom modules with custom message types, the Cosmos Indexer will not be able to parse and index these messages out of the box. The Cosmos Indexer will need to be extended to handle these custom modules and their message types.

## The Solution

The Cosmos Indexer Modules provide a set of Go packages that extend the Cosmos Indexer application to handle custom modules and their message types. These modules can be added to the Cosmos Indexer application to provide support for custom modules on a blockchain.

Each submodule in the Cosmos Indexer Modules package is designed to handle a either a specific custom module and its message types or the entire set of custom modules on a blockchain. The submodules provide the necessary Transaction Msg types and their, which can be registered with the Cosmos Indexer to parse and allow indexing transactions that contain these custom messages.

## Contributing

We welcome contributions to the Cosmos Indexer Modules. If you would like to contribute, please fork the repository and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

Contributions to add support for new custom modules are especially welcome. If you would like to add support for a custom module that is not currently supported, please open an issue to discuss the module and its message types, and we can work together to add support for it.
