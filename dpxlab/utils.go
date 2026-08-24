package dpxlab

// func nehmen => chan error ausgeben
func watch(f func() error) chan error {
	var errchan = make(chan error, 1)
	go func() {
		errchan <- f()
	}()

	return errchan
}
