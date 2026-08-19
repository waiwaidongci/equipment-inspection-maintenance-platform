package evidencebatch

import "errors"

type Batch struct {
	repo    Repository
	service Service
	opened  []*Resource
}

func NewBatch(repo Repository) *Batch { return &Batch{repo: repo, service: Service{}} }

func (b *Batch) RunBatch(values []string) error {
	var batchErr error
	for _, value := range values {
		resource := &Resource{}
		b.opened = append(b.opened, resource)
		if err := b.processOne(resource, value); err != nil {
			batchErr = errors.Join(batchErr, err)
		}
	}
	return errors.Join(batchErr, b.repo.Commit())
}

func (b *Batch) processOne(resource *Resource, value string) error {
	defer resource.Close()
	if err := b.service.Validate(value); err != nil {
		return err
	}
	return b.repo.Save(value)
}

func (b *Batch) Resources() []*Resource { return b.opened }
