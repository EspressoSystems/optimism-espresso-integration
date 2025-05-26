package benchmark

import (
	"math"
	"math/big"
	"slices"
	"time"

	geth_types "github.com/ethereum/go-ethereum/core/types"
)

// SingleL2TransactionMetric is a struct that holds metric information for a
// single transaction submitted to the L2 Sequencer.
// It is meant to hold information about the transaction's lifecycle, and
// progression through the system and it's various milestones.
//
// This information can then be aggregated and used to calculate and compare
// the performance of the system under various confirmations and load
// conditions.
type SingleL2TransactionMetric struct {
	SignedTransaction *geth_types.Transaction
	Receipt           *geth_types.Receipt

	LocalCreated   time.Time
	LocalSubmitted time.Time
	LocalReceipt   time.Time
	SeqReceipt     time.Time
	CaffReceipt    time.Time
	VerifyReceipt  time.Time
}

// Convert to Tracked Transactions
func (s *BenchmarkStats) IndividualTransactionMetrics() []SingleL2TransactionMetric {
	transactions := make([]SingleL2TransactionMetric, 0, len(s.Created))

	// Grab all of the transactions that were created, and track them
	for txHash, tx := range s.Created {
		value := SingleL2TransactionMetric{
			SignedTransaction: tx.Value,
			LocalCreated:      tx.Timestamp,
		}

		if submitted, ok := s.Submitted[txHash]; ok {
			value.LocalSubmitted = submitted.Timestamp
		}

		if receipt, ok := s.Receipts[txHash]; ok {
			value.Receipt = receipt.Value
			value.LocalReceipt = receipt.Timestamp
		}

		receipt := value.Receipt
		if receipt == nil {
			transactions = append(transactions, value)
			continue
		}

		if seqReceipt, ok := s.SeqReceipts[receipt.BlockHash]; ok {
			value.SeqReceipt = seqReceipt.Timestamp
		}

		if caffReceipt, ok := s.CaffReceipts[receipt.BlockHash]; ok {
			value.CaffReceipt = caffReceipt.Timestamp
		}

		if verifyReceipt, ok := s.VerifyReceipts[receipt.BlockHash]; ok {
			value.VerifyReceipt = verifyReceipt.Timestamp
		}

		transactions = append(transactions, value)
	}

	return transactions
}

// SplitIndividualTransactionMetrics is a function that takes a slice of
// SingleL2TransactionMetric and splits it into two slices: one for
// complete transactions and one for incomplete transactions. A
// transaction is considered complete if it has a receipt and a
// caff receipt. Otherwise, it is considered incomplete.
func SplitIndividualTransactionMetrics(
	transactions []SingleL2TransactionMetric,
) (complete, incomplete []SingleL2TransactionMetric) {
	complete = make([]SingleL2TransactionMetric, 0, len(transactions))
	incomplete = make([]SingleL2TransactionMetric, 0, len(transactions))
	var zeroTime time.Time

	for _, tx := range transactions {
		if tx.Receipt != nil && tx.CaffReceipt != zeroTime && tx.VerifyReceipt != zeroTime {
			complete = append(complete, tx)
		} else {
			incomplete = append(incomplete, tx)
		}
	}

	return complete, incomplete
}

// integer represents all of the integer types in Go.
type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// SampleSummary is a struct that holds the summary statistics of the given
// samples.
// It includes various measurements that are useful for understanding the
// distribution of the samples, and several key values.
type SampleSummary[T integer] struct {
	Count  int
	Min    T
	Max    T
	Mean   T
	Median T
	P99    T
	P90    T
	P75    T
	P50    T
	P25    T
	P10    T
	P01    T
	StdDev T
}

// SummarySamples is a function that takes a slice of sample values and then
// returns a SampleSummary struct that contains the summary statistics.
func SummarizeSamples[T integer](samples []T) SampleSummary[T] {
	if len(samples) <= 0 {
		return SampleSummary[T]{}
	}

	// Sort the samples
	slices.Sort(samples)

	l := len(samples)
	p1Index := l * 1 / 100
	p10Index := l * 10 / 100
	p25Index := l * 25 / 100
	p50Index := l * 50 / 100
	p75Index := l * 75 / 100
	p90Index := l * 90 / 100
	p99Index := l * 99 / 100

	metric := SampleSummary[T]{
		Count:  len(samples),
		Min:    samples[0],
		Max:    samples[l-1],
		Median: samples[p50Index],
		P99:    samples[p99Index],
		P90:    samples[p90Index],
		P75:    samples[p75Index],
		P50:    samples[p50Index],
		P25:    samples[p25Index],
		P10:    samples[p10Index],
		P01:    samples[p1Index],
	}

	total := new(big.Int)
	for _, sample := range samples {
		total.Add(total, big.NewInt(int64(sample)))
	}
	mean := new(big.Int).Div(total, big.NewInt(int64(len(samples))))
	metric.Mean = T(mean.Int64())

	// Calculate the standard deviation
	variance := new(big.Int)
	for _, duration := range samples {
		v := duration - metric.Mean
		variance.Add(variance, big.NewInt(int64(v*v)))
	}
	variance = variance.Div(variance, big.NewInt(int64(len(samples))))
	metric.StdDev = T(math.Sqrt(float64(variance.Int64())))

	return metric
}

// TimingMetrics is a struct that holds the timing metrics for various stages
// of the transaction lifecycle.
// It is currently populated with only a few timing references, but can be
// expanded to include more as necessary.
type TimingMetrics struct {
	CreatedToSubmitted      SampleSummary[time.Duration]
	SubmittedToLocalReceipt SampleSummary[time.Duration]
	SubmittedToReceipt      SampleSummary[time.Duration]
	ReceiptToCaff           SampleSummary[time.Duration]
	ReceiptToVerify         SampleSummary[time.Duration]
}

// ComputeCompletedTransactionStatistics is a function that takes a
// slice of SingleL2TransactionMetric and computes the statistics for
// the completed transactions. It returns a TimingMetrics struct that
// contains the summary statistics for the completed transactions.
//
// NOTE: This is tracking statistics on a transaction level. So it will
// exclude blocks that don't have any of the submitted transactions within
// them.
func ComputeCompletedTransactionStatistics(completed []SingleL2TransactionMetric) TimingMetrics {
	var zeroTime time.Time
	var createdToSubmittedSamples []time.Duration
	var submittedToReceiptSamples []time.Duration
	var submittedToLocalReceiptSamples []time.Duration
	var receiptToCaffSamples []time.Duration
	var receiptToVerifySamples []time.Duration

	for _, tx := range completed {
		if tx.LocalCreated == zeroTime {
			continue
		}

		if tx.LocalSubmitted == zeroTime {
			continue
		}

		createdToSubmittedSample := tx.LocalReceipt.Sub(tx.LocalSubmitted)
		createdToSubmittedSamples = append(createdToSubmittedSamples, createdToSubmittedSample)

		if tx.LocalReceipt == zeroTime {
			continue
		}

		submittedToReceiptSample := tx.LocalReceipt.Sub(tx.LocalSubmitted)
		submittedToReceiptSamples = append(submittedToReceiptSamples, submittedToReceiptSample)

		if tx.SeqReceipt != zeroTime {
			submittedToLocalReceiptSample := tx.SeqReceipt.Sub(tx.LocalSubmitted)
			submittedToLocalReceiptSamples = append(submittedToLocalReceiptSamples, submittedToLocalReceiptSample)
		}

		if tx.CaffReceipt != zeroTime {
			receiptToCaffSample := tx.CaffReceipt.Sub(tx.SeqReceipt)
			receiptToCaffSamples = append(receiptToCaffSamples, receiptToCaffSample)
		}

		if tx.VerifyReceipt != zeroTime {
			receiptToVerifySample := tx.VerifyReceipt.Sub(tx.SeqReceipt)
			receiptToVerifySamples = append(receiptToVerifySamples, receiptToVerifySample)
		}
	}

	return TimingMetrics{
		CreatedToSubmitted:      SummarizeSamples(createdToSubmittedSamples),
		SubmittedToReceipt:      SummarizeSamples(submittedToReceiptSamples),
		SubmittedToLocalReceipt: SummarizeSamples(submittedToLocalReceiptSamples),
		ReceiptToCaff:           SummarizeSamples(receiptToCaffSamples),
		ReceiptToVerify:         SummarizeSamples(receiptToVerifySamples),
	}
}

// ComputeRawMetricsStatistics is a function that takes a BenchmarkStats
// struct and computes the statistics for the transactions. It returns a
// TimingMetrics struct that contains the summary statistics for the
// transactions, and the blocks.
//
// NOTE: This attempts to track all references.  Any transaction based metric
// will be a necessarily be on a transaction level.  However, the sequencer and
// verifier metrics are on the block level.  This should allow for the tracking
// of blocks that do not have any of the submitted transactions within them.
// But it also means that the there will be fewer samples that are weight based
// on the number of transactions within the block.
func ComputeRawMetricsStatistics(stats BenchmarkStats) TimingMetrics {
	var createdToSubmittedSamples []time.Duration
	var submittedToReceiptSamples []time.Duration
	var submittedToLocalReceiptSamples []time.Duration
	var receiptToCaffSamples []time.Duration
	var receiptToVerifySamples []time.Duration

	// Inspect the transactions
	for txHash, created := range stats.Created {
		submitted, ok := stats.Submitted[txHash]
		if !ok {
			continue
		}

		createdToSubmittedSamples = append(createdToSubmittedSamples, submitted.Timestamp.Sub(created.Timestamp))
	}

	// This is specific Transaction receipts
	for txHash, localReceipt := range stats.Receipts {
		submitted, ok := stats.Submitted[txHash]
		if !ok {
			continue
		}

		submittedToLocalReceiptSamples = append(submittedToReceiptSamples, localReceipt.Timestamp.Sub(submitted.Timestamp))
		seqReceipt, ok := stats.SeqReceipts[localReceipt.Value.BlockHash]
		if !ok {
			continue
		}

		submittedToReceiptSamples = append(submittedToReceiptSamples, seqReceipt.Timestamp.Sub(submitted.Timestamp))
	}

	for blockHash, caffReceipt := range stats.CaffReceipts {
		seqReceipt, ok := stats.SeqReceipts[blockHash]
		if !ok {
			continue
		}

		receiptToCaffSamples = append(receiptToCaffSamples, caffReceipt.Timestamp.Sub(seqReceipt.Timestamp))
	}

	for blockHash, verifyReceipt := range stats.VerifyReceipts {
		seqReceipt, ok := stats.SeqReceipts[blockHash]
		if !ok {
			continue
		}

		receiptToVerifySamples = append(receiptToVerifySamples, verifyReceipt.Timestamp.Sub(seqReceipt.Timestamp))
	}

	return TimingMetrics{
		CreatedToSubmitted:      SummarizeSamples(createdToSubmittedSamples),
		SubmittedToLocalReceipt: SummarizeSamples(submittedToLocalReceiptSamples),
		SubmittedToReceipt:      SummarizeSamples(submittedToReceiptSamples),
		ReceiptToCaff:           SummarizeSamples(receiptToCaffSamples),
		ReceiptToVerify:         SummarizeSamples(receiptToVerifySamples),
	}
}
