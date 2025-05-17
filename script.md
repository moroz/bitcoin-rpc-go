In this video I'm going to show you how to quickly and securely generate disposable Bitcoin addresses for use in a payment processing system using the Go programming language.
We will use a method of key derivation defined in a document named "Bitcoin Improvement Proposal 32: Hierarchical Deterministic Wallets", or BIP32 in short.
You can use this method for online transactions, such as for sales on a website or in a mobile app, but with a little creativity you could also build a crypto-based point of sale system.

If you're new to this channel: My name is Karol Moroz, and you are watching Make Programming Fun Again.
I specialize in highly technical videos about programming.
This is my first video about Bitcoin, and hopefully a lot more are coming.

All the code shown in this video is available on my Github, and I will put a link to the repository in the description below.
I am recording this video on macOS Sequoia 15.4.1, but I will also test the code on Debian, Windows 11, and FreeBSD.

The first ingredient of our special crypto sauce is a "seed".
This "seed" is going to be your most cherished secret, the base of all wallets.
This seed must be 32 bytes long and be very difficult to guess, even with a supercomputer.
In computer parlance, we say that the seed must have "high entropy".

There is a bunch of ways you can generate a cryptographically secure seed, but on UNIX-like operating systems, you can use a tool called `openssl`.
When you run `openssl rand -base64 32`, you should get 32 random bytes, encoded in Base64.
If you are rawdogging a Windows development environment without WSL2, there is a PowerShell script in the repository just for you.

Next, we're going to have to store this seed in a safe place, to make sure noone else can read it, and that you're not going to lose it.
There are dozens of companies out there that will happily sell you a "cold wallet" for a hundred bucks, but unless you are currently targeted by Mosad, the IRS, or the FBI, a password manager like KeePassXC will suffice.

<!-- On Debian and Ubuntu, you can install KeePassXC using Apt, the package is named `keepassxc`, all lowercase. -->
<!-- On FreeBSD, there is a package with the same name, and you can install it using `pkg install`. -->
<!-- On macOS, you can install KeePassXC using Homebrew, using `brew install keepassxc`. -->
<!-- You can download the Windows installer from the official website of KeePassXC, but in order for it to run, you will also need to install a Microsoft-branded software library called "MSVC Redistributable", short for Microsoft Visual C++ Redistributable. -->

Next, I'm going to show you how to write a shell script that will generate a seed and store it in a brand new KeePassXC password database.
Before we start writing, though, you need to make sure that `keepassxc-cli` is executable.
I am not going to cover OS-specific installation instructions for KeePassXC in this video, but there's going to be a section in the README file in the source repository.
<!-- On Debian and FreeBSD, this should work just fine out of the box, and it should also be fine on macOS if you have installed KeePassXC using Homebrew. -->
<!-- If you are using WSL2 on Windows, the easiest way to get this script running is to install KeePassXC both natively on Windows and inside WSL2. -->
<!-- If you are rawdogging a Windows environment without WSL2, you can use my PowerShell script instead, I'll include it in the repository. -->

In the terminal, try running `keepassxc-cli`, and if you're getting an error message like "command not found", it means that `keepassxc-cli` is not in your `PATH`.

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

