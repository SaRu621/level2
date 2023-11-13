package pkg

type Builder struct {
	Collector Collector
}

func NewBuilder(collector Collector) *Builder {
	return &Builder{Collector: collector}
}

func (builder *Builder) SetCollector(collector Collector) {
	builder.Collector = collector
}

func (builder Builder) CreateComputer() Computer {
	builder.Collector.SetCore()
	builder.Collector.SetBrand()
	builder.Collector.SetMemory()
	builder.Collector.SetGraphicCard()
	builder.Collector.SetMonitor()

	return builder.Collector.GetComputer()
}
