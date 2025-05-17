In this video I'm going to show you how to quickly and securely generate disposable Bitcoin addresses for use in a payment processing system using the Go programming language.
We will use a method of key derivation defined in a document named "Bitcoin Improvement Proposal 32: Hierarchical Deterministic Wallets", or BIP32 in short.
You can use this method for online transactions, such as for sales on a website or in a mobile app, but with a little creativity you could also build a crypto-based point of sale system.

If you're new to this channel: My name is Karol Moroz, and you are watching Make Programming Fun Again.
I specialize in highly technical videos about programming.
This is my first video about Bitcoin, and hopefully a lot more are coming.

Let's start by generating a random, 32-byte-long seed. This is the base, high-entropy secret on which all private keys will be based.
There is a bunch of ways you can generate a seed, but on UNIX-based operating systems, you can use a tool called `openssl`.
When you run `openssl rand -base64 32`, you should get 32 random bytes, encoded in Base64.

Next, we're going to have to store this seed in a safe place, to make sure noone else can read it, and that you're not going to lose it.
Now, there are dozens of companies out there that will happily sell you a "cold wallet" for a hundred bucks, but unless you are currently targeted by Mosad, the IRS, or the FBI, a password manager like KeePassXC will suffice.

Using BIP32, you can generate a virtually infinite number of public keys and addresses from a single 32-byte seed.
Now, the best part is that you can generate public keys based on a single public key, without ever storing the seed on your server.

Before I show you exactly how to do it, let us quickly discuss the rationale.
Why should you generate burner addresses and not just accept payments on your hot wallet or to your crypto exchange?

The first reason is privacy. All transactions on the Bitcoin blockchain are public, so if you use the same address for everything, anyone who has ever bought from you will be able to tell exactly how much money you have received and spent. That is not ideal.

The second reason is technical.
If you assign a unique Bitcoin address to each payment, then whenever funds arrive at that address, you can unambiguously link them to a specific transaction.
Combine that with a well-designed payment processing system, and you can build a website that practically prints money on autopilot.

Once you have processed a number of payments, you can move the funds from several disposable addresses to your wallet in a process known as *sweeping*.
In order to do that, you're going to need the private keys for all wallets with funds, and for that purpose, you're going to need the original 32-byte seed, as well as the derivation paths for all addresses.

---

Now, without further ado, let's write some code!

