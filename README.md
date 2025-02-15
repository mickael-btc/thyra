# thyra

[![CI](https://github.com/massalabs/thyra/actions/workflows/CI.yml/badge.svg)](https://github.com/massalabs/thyra/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/massalabs/thyra/branch/main/graph/badge.svg?token=592LPZLC4M)](https://codecov.io/gh/massalabs/thyra)
[![Go Report Card](https://goreportcard.com/badge/github.com/massalabs/thyra)](https://goreportcard.com/report/github.com/massalabs/thyra)

An entrance to the Massa blockchain.

## ⚠️ WIP

This project is still WIP. It is a prototype. 

⚠️ Potential breaking changes ahead ⚠️

## Contribute
go to [Contributing](./CONTRIBUTING.md)
## How to...

### ... install Thyra on my computer?

Follow the instructions for your computer in the wiki:

- [MacOS](https://github.com/massalabs/thyra/blob/main/INSTALLATION.md#macos)
- [Linux](https://github.com/massalabs/thyra/blob/main/INSTALLATION.md#linux)
- [Windows](https://github.com/massalabs/thyra/blob/main/INSTALLATION.md#windows)

### ... manage my wallet?
1. Create / delete your wallet 

You can access to Thyra wallet interface at URL : <https://my.massa/thyra/wallet/index.html>

By inputing the 'Nickname' & 'Password', you'll be able to create an encrypted wallet locally on your machine.
To delete your wallet, simply use the interface. 
⚠️ If you delete your wallet, you won't be able to edit the website linked to it anymore.

2. Get coins on your wallet

To get coins on your wallet, you have to send your address on [Massa faucet channel](https://discord.com/channels/828270821042159636/866190913030193172)
Make sure that you use the latest version of Thyra (and defacto Testnet) [here](https://github.com/massalabs/thyra/releases/latest/), otherwise the faucet won't work.
NB: if you just installed Thyra, you're using the latest version!

### ... store a website on chain?

You can access to Thyra web hosting interface at URL : <http://my.massa/thyra/websiteCreator/index.html>

In order to register a website on Thyra you'll need to :

- Deploy a Smart Contract that will handle the storage of your website, your DNS name will fetch the Address of this Smart Contract - this step is handled by Thyra, don't worry. But at least you know what's happening behind the wood.
- Upload the build of your application.
  * If you don't have an application or website handy but still want to test this functionality, here are some websites you can use:
    - [bbx.zip](https://github.com/massalabs/thyra/files/10169142/bbx.zip)
    - [flappy.zip](https://github.com/massalabs/thyra/files/10169143/flappy.zip)
    - [tic-tac-toe.zip](https://github.com/massalabs/thyra/files/10169144/tic-tac-toe.zip)

- Use a wallet with sufficient coins to upload it on the blockchain
Important note: At the moment, we have defined that 1 chunk (=280ko) of data worth 100 MAS. It will change and become more and more specific and precise as the Testnet and Thyra are evolving. In the mean time, we have defined it arbitrarily.
- Visit the registry to find your .massa website and see others

**Troubleshooting** If you got any problem at any steps of the process, join us on [Discord dedicated channel](https://discord.com/channels/828270821042159636/851942484212318259) or [report a problem](https://github.com/massalabs/thyra/issues/new/choose)



### ... pass options to `thyra-server`?

Thyra accepts different options that you can specify when you start the program.
In this section you will find a non-exhaustive list of such options and examples of how you can use them.

--node-server : Specify which Massa network Thyra will communicate with while running.
Accepts a URL, an IP address or one of the following values :

- TESTNET : Uses Massa's testnet
- LABNET : Uses Massa's labnet
- INNONET : Uses Massa's innonet
- LOCALHOST : Expect Massa's network to be hosted at 127.0.0.1

To use this option with a constant, you have to execute :
`thyra-server --node-server=LABNET`
To use this option with a custom IP address, you have to execute :
`thyra-server --node-server=192.168.X.X`

### ... code with auto-reload

You can run the application with this command: `air`.

It will generate and reload thyra each time a file is modified.

## Usage

### Upload a website

With Thyra you can upload a website. It will be hosted on the Massa blockchain and available as a `.massa` URL. You can access this URL through Tyhra.

For this to work, the file you upload must be a zip archive (file ending with `.zip`). This archive must contain a `index.html` file at the root.

## Additional information

### Why this name?

θύρα (thýra) in ancient Greek means door, entrance. This is exactly what this project is: an entrance to the Massa blockchain.

### How to pronounce it?

See <https://www.youtube.com/watch?v=_0BQ7sSJMTw>.
