package key_service

func (k_svc *KeyService) List() ([]string, error) {
	keys, err := k_svc.storage.List()
	if err != nil {
		return nil, err
	}

	return []string(keys), nil
}