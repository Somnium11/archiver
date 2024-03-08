package cmd

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "packfile using variable-lenght code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vlc called")
	},
}

	const packedExtension = "vlc"
	var ErrEmptyPath = errors.New("path to file is not specified")
func pack(_ *cobra.Command, args []string) {
if len(args) == 0 || args[0] == "" {
	handleErr(errors.New("no file specified"))
}
	filepath := args[0]

	r, err :=os.Open(filepath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()
	
	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}
// packed := Encode(data)
packed := ""
fmt.Println(string(data))

err := os.WriteFile(packedFileName, []byte(packed), 0644)
if err != nil {
	handleErr(err)
}
}

func packedFileName(paht string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, ext := filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(vlcCmd)
}