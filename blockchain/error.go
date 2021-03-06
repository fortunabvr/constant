// Copyright (c) 2014-2016 The thaibaoautonomous developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain

import "fmt"

const (
	UnExpectedError = iota
	UpdateMerkleTreeForBlockError
	UnmashallJsonBlockError
	CanNotCheckDoubleSpendError
)

var ErrCodeMessage = map[int]struct {
	code    int
	message string
}{
	UnExpectedError:               {-1, "Unexpected error"},
	UpdateMerkleTreeForBlockError: {-2, "Update Merkle Commitments Tree For Block is failed"},
	UnmashallJsonBlockError:       {-3, "Unmarshall json block is failed"},
	CanNotCheckDoubleSpendError:   {-4, "Unmarshall json block is failed"},
}

type BlockChainError struct {
	Code    int
	Message string
	Err     error
}

func (e BlockChainError) Error() string {
	return fmt.Sprintf("%d: %s \n %+v", e.Code, e.Message, e.Err)
}

func NewBlockChainError(key int, err error) *BlockChainError {
	return &BlockChainError{
		Code:    ErrCodeMessage[key].code,
		Message: ErrCodeMessage[key].message,
		Err:     err,
	}
}
