package main

type FinanceReport struct {
}

type Report struct {
}

func (r *FinanceReport) MakeReport() *Report {
	return &Report{}
}

func main() {

}
