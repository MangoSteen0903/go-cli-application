package util

func HandleErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
