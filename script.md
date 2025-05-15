In this video I'm going to show you how to quickly and securely generate disposable Bitcoin addresses for use in a payment processing system using the Go programming language.
We will be utilizing a method of key derivation defined in a document named "Bitcoin Improvement Proposal 32: Hierarchical Deterministic Wallets", or BIP32 in short.
You can use this method for online transactions, such as for sales on a website or in a mobile app, but with a little creativity you could also build a crypto-based point of sale system.

If you're new to this channel: My name is Karol Moroz, and you are watching Make Programming Fun Again.
I specialize in highly technical videos about programming.
This is my first video about Bitcoin, and hopefully a lot more are comming.

Before we talk about the "hows", let's quickly talk about the "whys". Why should you generate burner addresses and not just accept payments on your hot wallet or to your crypto exchange?

The first reason is privacy. All transactions on the Bitcoin blockchain are public, so if you use the same address for everything, anyone who has ever bought from you will be able to tell exactly how much money you have received and spent. That is not ideal.

The second reason is technical.
If you assign a unique Bitcoin address to each payment, then whenever funds arrive at that address, you can unambiguously link them to a specific transaction.
Combine that with a well designed payment processing system, and you can build a Website that practically prints money on autopilot.

<!-- One downside of this approach is that, with many fragmented wallets and addresses, it can become difficult to spend the funds you have received. -->
<!-- You can solve this by consolidating everything into a single wallet--a process known as sweeping, which I will cover in another video. -->

---

Now, without further ado, let's write some code!

