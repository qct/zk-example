package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"os"
)

var (
	pkPath       = "zk-example.pk"
	vkPath       = "zk-example.vk"
	contractPath = "solidity/contracts/zk-example.sol"
	witnessPath  = "solidity/witness.json"
	proofPath    = "solidity/proof"
)

type MyCircuit struct {
	X frontend.Variable `gnark:"x"`       // x  --> secret visibility (default)
	Y frontend.Variable `gnark:",public"` // Y  --> public visibility
}

// Define y = x^4 + x^3 + x^2 + x
func (c *MyCircuit) Define(api frontend.API) error {
	s1 := api.Mul(c.X, c.X)
	s2 := api.Mul(s1, c.X)

	result := api.Add(c.X, s1, s2, api.Mul(s2, c.X))
	api.AssertIsEqual(c.Y, result)
	return nil
}

func main() {
	// compiles our circuit into a R1CS
	var circuit MyCircuit
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	pk, vk, err := generatePkVkIfNotExist(err, r1cs)

	// create a valid proof
	assignment := MyCircuit{X: 1, Y: 4}
	// witness creation
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	// export witness for understanding
	publicWitness, _ := witness.Public()
	err = exportWitness(publicWitness, &circuit)
	if err != nil {
		panic(err)
	}

	// groth16: Prove & Verify
	proof, _ := groth16.Prove(r1cs, pk, witness)
	err = exportProof(proof)
	if err != nil {
		panic(err)
	}

	// ensure gnark (Go) code verifies it
	publicWitness, _ = witness.Public()
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		panic(err)
	}
}

func generatePkVkIfNotExist(err error, r1cs constraint.ConstraintSystem) (groth16.ProvingKey, groth16.VerifyingKey, error) {
	root, _ := os.Getwd()
	_, err = os.Stat(root + string(os.PathSeparator) + pkPath)
	_, err = os.Stat(root + string(os.PathSeparator) + vkPath)
	if err != nil {
		fmt.Printf("Pk or Vk does not exist, will generate a new pair\n")
		// groth16 zkSNARK: Setup
		pk, vk, err := groth16.Setup(r1cs)
		if err != nil {
			panic(err)
		}

		// export PK
		err = exportPK(pk)
		if err != nil {
			panic(err)
		}
		// export VK
		err = exportVK(vk)
		if err != nil {
			panic(err)
		}
		//export contracts
		err = exportContract(vk)
		if err != nil {
			panic(err)
		}
		return pk, vk, err
	} else {
		pk := groth16.NewProvingKey(ecc.BN254)
		pkFile, err := os.Open(root + string(os.PathSeparator) + pkPath)
		if err != nil {
			panic("Pk does not exist")
		}
		_, err = pk.ReadFrom(pkFile)
		_ = pkFile.Close()

		vk := groth16.NewVerifyingKey(ecc.BN254)
		vkFile, err := os.Open(root + string(os.PathSeparator) + vkPath)
		if err != nil {
			panic("Vk does not exist")
		}
		_, err = vk.ReadFrom(vkFile)
		_ = vkFile.Close()
		return pk, vk, err
	}
}

func exportProof(proof groth16.Proof) error {
	root, _ := os.Getwd()
	f, err := os.Create(root + string(os.PathSeparator) + proofPath)
	if err != nil {
		return err
	}
	_, err = proof.WriteRawTo(f)
	return err
}

func exportWitness(w witness.Witness, circuit *MyCircuit) error {
	root, _ := os.Getwd()
	f, err := os.Create(root + string(os.PathSeparator) + witnessPath)
	if err != nil {
		return err
	}
	schema, err := frontend.NewSchema(circuit)
	if err != nil {
		return err
	}
	w.MarshalBinary()
	json, err := w.ToJSON(schema)
	if err != nil {
		return err
	}
	_, err = f.Write(json)
	return err
}

func exportContract(vk groth16.VerifyingKey) error {
	root, _ := os.Getwd()
	f, err := os.Create(root + string(os.PathSeparator) + contractPath)
	if err != nil {
		return err
	}
	err = vk.ExportSolidity(f)
	return err
}

func exportVK(vk groth16.VerifyingKey) error {
	root, _ := os.Getwd()
	f, err := os.Create(root + string(os.PathSeparator) + vkPath)
	if err != nil {
		return err
	}
	_, err = vk.WriteRawTo(f)
	return err
}

func exportPK(pk groth16.ProvingKey) error {
	root, _ := os.Getwd()
	f, err := os.Create(root + string(os.PathSeparator) + pkPath)
	if err != nil {
		return err
	}
	_, err = pk.WriteRawTo(f)
	return err
}
