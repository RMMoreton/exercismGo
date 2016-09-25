// Package diffiehellman solves an Exercism challenge.
package diffiehellman

import (
	"math/big"
	"math/rand"
)

// PrivateKey returns a random *big.Int that's less than the passed
// *big.Int, and greater than 1.
func PrivateKey(p *big.Int) *big.Int {
	toReturn := big.NewInt(0)
	bitsNeeded := p.BitLen()
	for ; bitsNeeded > 0; bitsNeeded -= 63 {
		toReturn.Or(toReturn, big.NewInt(rand.Int63()))
		toReturn.Lsh(toReturn, 63)
	}
	toReturn.Rsh(toReturn, 63+uint(-bitsNeeded))
	// If my random integer is still too big, shift it right one more time.
	if toReturn.Cmp(p) != -1 {
		toReturn.Rsh(toReturn, 1)
	}
	// If my random integer is less than 2, set it equal to 2.
	// This *probably* won't happen, except when p is small.
	if toReturn.Cmp(big.NewInt(2)) == -1 {
		toReturn = big.NewInt(2)
	}
	return toReturn
}

// PublicKey takes a private key and two primes and returns a
// public key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	toReturn := big.NewInt(0)
	bigG := big.NewInt(0)
	bigG.SetInt64(g)
	toReturn.Exp(bigG, private, p)
	return toReturn
}

// NewPair takes two primes and returns a public key and
// a private key based on those two primes.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return private, public
}

// SecretKey takes a prime, a private key, and a public key
// and returns a secret.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	toReturn := big.NewInt(0)
	toReturn.Exp(public2, private1, p)
	return toReturn
}
