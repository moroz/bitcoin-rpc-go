package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

func generateSeed() []byte {
	var buf = make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		log.Fatal(err)
	}
	return buf
}

func locateKeepassXC() (string, error) {
	if bin, err := exec.LookPath("keepassxc-cli"); err == nil {
		return bin, nil
	}
	if runtime.GOOS == "windows" {
		return exec.LookPath(`C:\Program Files\KeePassXC\keepassxc-cli`)
	}
	if runtime.GOOS == "darwin" {
		return exec.LookPath("/Applications/KeePassXC/Contents/MacOS/keepassxc-cli")
	}
	return "", errors.New("keepassxc-cli not found in PATH")
}

const FILENAME = "btc-seed.kdbx"

// resolveDatabasePath returns a path where the database will be created.
// If the program is run inside a Git repository, the database will be created in its parent directory.
// If no Git directory is found, the database will be created in the user's home directory.
func resolveDatabasePath() (string, error) {
	// Limit search to $HOME
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Start searching from PWD
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		gitDir := filepath.Join(dir, ".git")

		// If a git repository exists, return one level above git root
		if _, err := os.Stat(gitDir); err == nil {
			return filepath.Join(filepath.Dir(dir), FILENAME), nil
		}

		// If we hit root path, create the database in the home directory
		if filepath.Dir(dir) == dir {
			return filepath.Join(home, FILENAME), nil
		}

		// Go one level up
		dir = filepath.Dir(dir)
	}
}

func promptPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func createDB(path, passphrase string) error {
	bin, err := locateKeepassXC()
	if err != nil {
		return err
	}
	outBuf := bytes.NewBuffer([]byte{})
	cmd := exec.Command(bin, "db-create", "-q", "-p", "-t 3000", path)
	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("%s\n%s\n", passphrase, passphrase))
	cmd.Stdout = outBuf
	cmd.Stderr = outBuf
	return cmd.Run()
}

func storeSeedInDB(path, passphrase string, seed []byte) error {
	bin, err := locateKeepassXC()
	if err != nil {
		return err
	}
	outBuf := bytes.NewBuffer([]byte{})
	cmd := exec.Command(bin, "add", path, "/MASTER_SEED", "-p")
	seedB64 := base64.StdEncoding.EncodeToString(seed)
	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("%s\n%s\n", passphrase, seedB64))
	cmd.Stdout = outBuf
	cmd.Stderr = outBuf
	return cmd.Run()
}

func main() {
	_, err := locateKeepassXC()
	if err != nil {
		log.Fatal(err)
	}

	dbPath, err := resolveDatabasePath()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("This script will create a KeePassXC database at %s and store a master seed in this database.\n", dbPath)

	if _, err := os.Stat(dbPath); err == nil {
		log.Fatalf("File %s exists. Please delete this file to proceed.", dbPath)
	}

	passphrase, err := promptPassword("Please define a passphrase: ")
	if err != nil {
		log.Fatal(err)
	}
	confirmation, err := promptPassword("Confirm passphrase: ")
	if err != nil {
		log.Fatal(err)
	}
	if passphrase != confirmation {
		fmt.Println("FATAL: Passphrase does not match confirmation!")
		os.Exit(1)
	}

	log.Println("Creating a KeePassXC database, this may take some time...")
	if err := createDB(dbPath, passphrase); err != nil {
		log.Fatal(err)
	}

	log.Println("Storing seed in database...")
	if err := storeSeedInDB(dbPath, passphrase, generateSeed()); err != nil {
		log.Fatal(err)
	}
}
