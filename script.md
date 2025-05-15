In this video I'm going to show you how to quickly and securely generate single-use Bitcoin addresses for use in a payment processing system.
You can use this method for online transactions, such as for sales on a website or in a mobile app, but with a little creativity you could also build a crypto-based point of sale system.

If you're new to this channel: My name is Karol Moroz, and you are watching Make Programming Fun Again.
I specialize in highly technical videos about programming.
This is my first video about Bitcoin, and hopefully a lot more are comming.

Before we talk about the "hows", let's quickly talk about the "whys". Why should you generate burner addresses and not just accept payments on your hot wallet or to your crypto exchange?

The first reason is privacy. All transactions on the Bitcoin blockchain are public, so if you use the same address for everything, anyone who has ever bought from you will be able to tell exactly how much money you have received and spent. That is not ideal.

The second reason is technical.
If you assign a unique Bitcoin address to each payment, then whenever funds arrive at that address, you can unambiguously link them to a specific transaction.
Combine that with a well designed payment processing system, and you can build a Website that practically prints money on autopilot.

When building a software system that accepts Bitcoin payments, it is usually a good idea to use a unique wallet address for each payment.
In this video, I'm going to show you how to quickly and securely generate disposable Bitcoin wallets using the Go programming language.

For the seller, this approach has multiple advantages.
First and foremost, it gives you privacy.
Since the records of all Bitcoin transactions are publicly available on the Blockchain, if you use the same address more than once, other people can very easily estimate how much money you have received in total.
On the other hand, if you use a new address every time you accept a payment, it becomes much more difficult for others to spy on your business.

Using disposable addresses is also convenient from the developer's perspective.
There is always exactly one address for each payment request, so when a payment arrives, it's very easy to find out what the payment is for.
you can very easily verify the payment against the order data in your database.

A downside of this approach is that with a lot of fragmented wallets and addresses, it becomes difficult to spend the money you have received. You can solve this by consolidating all that money in a single wallet. This approach is called sweeping, and I will cover it in another video.


