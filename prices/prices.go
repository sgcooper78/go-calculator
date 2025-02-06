package prices

import (
	"fmt"

	"github.com/sgcooper78/go-calculator/conversion"
	"github.com/sgcooper78/go-calculator/file_manager"
)

type TaxIncludedPriceJob struct {
	IOManager         file_manager.FileManager `json:"-"`
	TaxRate           float64                  `json:"tax_rate"`
	InputPrices       []float64                `json:"input_prices"`
	TaxIncludedPrices map[string]string        `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) loadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Error converting strings to floats:", err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	fmt.Println(result)

	err := job.IOManager.WriteResult(job)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func NewTaxIncludedPriceJob(fm file_manager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
