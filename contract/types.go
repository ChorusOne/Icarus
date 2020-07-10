package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type UpvotesMap map[string]*big.Int
type IndexMap map[string]*big.Int

type DetailedQueuedProposal struct {
	ProposalID  *big.Int       `json:"proposalID"`
	Index       *big.Int       `json:"index"`
	BlockNumber *big.Int       `json:"currentBlockNumber"`
	Stage       string         `json:"stage"`
	Proposer    common.Address `json:"proposer"`
	Upvotes     *big.Int       `json:"upvotes"`
	Description string         `json:"description"`

	ProposalEpoch   *big.Int `json:"proposalEpoch"`
	ReferendumEpoch *big.Int `json:"referendumEpoch"`
	ExecutionEpoch  *big.Int `json:"executionEpoch"`
	ExpirationEpoch *big.Int `json:"expirationEpoch"`

	QueuedAtBlockNumber *big.Int `json:"queuedAtBlockNumber,omitempty"`
	Deposit             *big.Int `json:"deposit"`
	QueuedTimestamp     *big.Int `json:"queuedAtTimestamp,omitempty"`
}

type DetailedProposal struct {
	ProposalID   *big.Int       `json:"proposalID"`
	Index        *big.Int       `json:"index"`
	BlockNumber  *big.Int       `json:"currentBlockNumber"`
	Stage        string         `json:"stage"`
	Proposer     common.Address `json:"proposer"`
	YesVotes     *big.Int       `json:"yesVotes,omitempty"`
	NoVotes      *big.Int       `json:"noVotes,omitempty"`
	AbstainVotes *big.Int       `json:"abstainVotes,omitempty"`
	Description  string         `json:"description"`

	ProposalEpoch   *big.Int `json:"proposalEpoch"`
	ReferendumEpoch *big.Int `json:"referendumEpoch"`
	ExecutionEpoch  *big.Int `json:"executionEpoch"`
	ExpirationEpoch *big.Int `json:"expirationEpoch"`

	QueuedAtBlockNumber *big.Int `json:"queuedAtBlockNumber,omitempty"`
	Deposit             *big.Int `json:"deposit"`
	QueuedTimestamp     *big.Int `json:"queuedAtTimestamp,omitempty"`
}

//DetailedExpiredProposal represents a detailed expired proposal (Proposal in Expiration Stage)
type DetailedExpiredProposal struct {
	ProposalID  *big.Int       `json:"proposalID"`
	BlockNumber *big.Int       `json:"currentBlockNumber"`
	Stage       string         `json:"stage"`
	Proposer    common.Address `json:"proposer"`
	Executed    bool           `json:"executed"`

	QueuedAtBlockNumber *big.Int `json:"queuedAtBlockNumber"`
	Deposit             *big.Int `json:"deposit"`
	QueuedTimestamp     *big.Int `json:"queuedAtTimestamp"`
	Description         string   `json:"description"`
}

//AssortedProposals is a collection of proposals across stages
type AssortedProposals struct {
	QueuedProposals     []*DetailedQueuedProposal  `json:"queuedProposals"`
	ApprovalProposals   []*DetailedProposal        `json:"approvalProposals"`
	ReferendumProposals []*DetailedProposal        `json:"referendumProposals"`
	ExecutionProposals  []*DetailedProposal        `json:"executionProposals"`
	ExpiredProposals    []*DetailedExpiredProposal `json:"expiredProposals"`
}
