package aws

func convertString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
