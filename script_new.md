In this video I'm going to show you how to generate single-use Bitcoin wallets using the Go programming language.
You can use this method in a payment processing system, for example to sell products and services or a website.

Start by creating a Go project.
Create a new directory to store the project.
I'm going to call this project "michi", which is the Japanese word for "road".
In this directory, initialize a Go module using `go mod init github.com/moroz/michi`.
`moroz` is my Github username, and `michi` is the name of the project.
Initialize an empty Git repository using `git init`.
Stage all files for commit using `git add .`. Create an inital commit using `git commit`.

Using `mkdir -p`, create a directory at `cmd/seed`, and `cd` into this directory.
Create an empty file called `main.go` and open it in a code editor.
Define `package main` and `func main()`.

First, let's generate a seed value.
Allocate a new byte slice with the length of 32.
Using `rand.Read`, fill the slice with cryptographically secure random bytes.
Next, let's print the value in Base64 to make sure this works.
When we run the program using `go run .`, it will print a random Base64-encoded string.
Note that this value is going to be different each time you run this program.
This is fine for now.

Next, install a Go library called `btcutil`.
It is a part of a larger project called `btcd`, which is a full Bitcoin node implementation in Go.
In the terminal, run `go get github.com/btcsuite/btcd/btcutil/hdkeychain`.
This is going to install a package called `hdkeychain`, responsible for generating "hierarchical deterministic wallets".

This may sound like a word salad, but what this means is that you can generate a number of key pairs from a single seed.
More specifically, you can generate over 4 billion derived key pairs from a single parent key pair.
These keys and addresses can be derived either using "hardened" derivation, which requires the parent's private key, or using non-hardened derivation, which only requires the parent's public key.
The exact algorithm used to derive these key pairs is defined in a document called "Bitcoin Improvement Proposal 32: Hierarchical Deterministic Wallets", or BIP32 for short.

You can repeat the derivation step multiple times to derive several levels of key pairs.
If you record the indices used for each derivation step, you get what's called a "derivation path".
In this notation, the master key is represented by a lowercase letter "m," and derivation levels are separated by slashes.
Hardened derivation is indicated either by a prime symbol, (a fancy mathematical term for an apostrophe), or the lowercase letter "h" (for "hardened").

The derivation algorithm is deterministic, which means that if you start with the same seed and follow the same derivation path, you will arrive at the same results.
In fact, this is exactly how popular wallet software manages multiple addresses within your wallet.

For the purposes of our project, we can take advantage of one particular property of this algorithm, namely that public keys and addresses can be derived from a public key.
This means that the application generating addresses does not need to know the secret seed, 

but what this means in practice is that you can derive 

some Go libraries to derive addresses from this seed.
We will use a method of key derivation defined in a document named BIP32, which is short for "Bitcoin Improvement Proposal 32: Hierarchical Deterministic Wallets".
This may sound like a word salad, but the "wallets" are really just pairs of private and public keys.
The private key proves that you own an address, but the public key is enough to receive a payment.

This is going to be the initial secret value from which all 

You can use this method for any type of online transactions, such as for sales on a website or in a mobile app, and with a little creativity you could even build a crypto-based point of sale system.

<!-- We need the private key to spend coins, but having just the public key is enough to calculate a Bitcoin address and to receive a payment. -->
We can call the wallets "hierarchical", because keys can be derived from a single secret in a tree-like structure, all based on the same seed value.
At the same time, 
<!-- "Deterministic" means that no matter how many times you perform the same operation, as long as the input values are the same, the result is going to be the same. -->
To sum up: We can generate pairs of private and public keys in a top-down structure, all based on the same seed value.
Now, for me the best part is that public keys can actually be derived from public keys, so you don't have to store the private key on the server.

If you're new to this channel: My name is Karol Moroz, and you are watching Make Programming Fun Again.
I specialize in highly technical videos about programming.
This is my first video about Bitcoin, so if I got something wrong, please let me know in the comments, I am here to learn.

There are dozens of companies out there that will happily sell you crypto wallets for a hundred dollars or more, but in its essence, a Bitcoin address 

When receiving Bitcoin payments in an online setting, 

Bitcoin addresses are often compared to bank accounts, because they can hold money
When receiving a Bitcoin payment
There are two reasons to use 
