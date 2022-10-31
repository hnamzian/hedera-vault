package key_service

func (k_svc *KeyService) Delete(id string) error {
	return k_svc.storage.Delete(id)
}