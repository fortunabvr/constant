package database

type CandidateElement struct {
	PubKey     []byte
	VoteAmount uint64
}

type CandidateList []CandidateElement

func (A CandidateList) Len() int {
	return len(A)
}
func (A CandidateList) Swap(i, j int) {
	A[i], A[j] = A[j], A[i]
}
func (A CandidateList) Less(i, j int) bool {
	return A[i].VoteAmount < A[j].VoteAmount
}
