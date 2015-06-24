package ase

import "testing"

// STraightforward to write, come back to this 
func TestASEStruct(t *testing.T) {
    t.SkipNow()
}

func TestDecode(t *testing.T) {
    
    // Load and decode the ase test file
    filePath := "testfiles/test.ase"
    ase := ASE{}
    ase.Decode(filePath, false)
    
    // Check that each respective field is correctly populated
    expectedSignature := "ASEF"
    if ase.Signature != expectedSignature {
        t.Error("expected signature of ASEF, got ", ase.Signature)   
    }
    
    expectedVersion := [2]int16{1, 0}
    if ase.Version != expectedVersion {
        t.Error("expected version of ", expectedSignature,
                " got ", ase.Signature)
    }
    
    expectedNumBlocks := [1]int32{10}
    if ase.NumBlocks != expectedNumBlocks {
        t.Error("expected NumBlocks of ", expectedNumBlocks,
                " got ", ase.NumBlocks)  
    }
}