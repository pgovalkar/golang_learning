package core

func FileOpertations(file_name string, service string) {
	if file_name == "" {
		file_name = "sub_orig_dashboard.json"
	}
	file, err := ReadFile(file_name)
	if err != nil {
		WriteFile(file_name, file, service)
	}
}
