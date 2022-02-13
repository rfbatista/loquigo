package templatepool

func NewTemplatePoolService() TemplatePoolService {
	return TemplatePoolService{}
}

type TemplatePoolService struct{}

func (t TemplatePoolService) Run() {}
